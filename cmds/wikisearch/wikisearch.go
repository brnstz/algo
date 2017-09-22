package main

import (
	"bufio"
	"compress/bzip2"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/brnstz/algo"
)

const (
	loadLogInterval  = 1000000
	maxCompletions   = 25
	dumpDate         = "20170820"
	wikiIndexURL     = "http://dumps.wikimedia.your.org/%vwiki/%v/%vwiki-%v-pages-articles-multistream-index.txt.bz2"
	streamURL        = "https://stream.wikimedia.org/v2/stream/recentchange"
	wikiCodes        = "en|ceb|sv|de|nl|fr|ru|it|es|war|pl|vi|ja|pt|zh|uk|fa|ca|ar|no|sh|fi|hu|id|ko"
	downloadWorkers  = 5
	titleField       = 3
	streamDataPrefix = "data: "
)

var writeLock sync.Mutex

type completion struct {
	Word  string   `json:"word"`
	Wikis []string `json:"wikis"`
}
type wordResponse struct {
	Exists      bool         `json:"exists"`
	Completions []completion `json:"completions"`
	Wikis       []string     `json:"wikis"`
}

type dlReq struct {
	wiki string
	mask int64
}

type wikiStream struct {
	Title      string `json:"title"`
	ServerName string `json:"server_name"`
}

func loadStream(masks map[string]int64, t *algo.Trie) {
	var ws wikiStream

	// Continue forever if we are disconnected
	for {
		func() {

			// Open up the stream URL
			resp, err := http.Get(streamURL)
			if err != nil {
				log.Printf("can't download %v: %v\n", streamURL, err)
				return
			}
			defer resp.Body.Close()

			// Read each line
			s := bufio.NewScanner(resp.Body)
			for s.Scan() {
				text := s.Text()

				// Check if we have the "data: " prefix
				if strings.HasPrefix(text, streamDataPrefix) {

					// Read the JSON data after the prefix until the
					// end of the line.
					jbytes := []byte(text[len(streamDataPrefix):len(text)])
					err := json.Unmarshal(jbytes, &ws)
					if err != nil {
						log.Printf("can't unmarshal: %v", err)
						continue
					}

					// Pick which wiki this is in, if it's not one
					// we support, then forget it
					wiki := strings.Split(ws.ServerName, ".")[0]
					mask, exists := masks[wiki]
					if !exists {
						continue
					}

					// Add to our trie
					writeLock.Lock()
					log.Println(ws.Title, mask, wiki)
					t.Add(ws.Title, mask)
					writeLock.Unlock()
				}
			}
		}()
	}
}

func download(reqs chan dlReq, t *algo.Trie) {

	for req := range reqs {
		url := fmt.Sprintf(wikiIndexURL, req.wiki, dumpDate, req.wiki, dumpDate)

		log.Printf("loading from %v.wikipedia.org", req.wiki)

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("can't download %v: %v\n", url, err)
			return
		}
		defer resp.Body.Close()

		s := bufio.NewScanner(bzip2.NewReader(resp.Body))
		i := 0
		for s.Scan() {
			parts := strings.SplitN(s.Text(), ":", titleField)
			if len(parts) == titleField {
				writeLock.Lock()
				t.Add(parts[titleField-1], req.mask)
				writeLock.Unlock()
				i += 1

				if i%loadLogInterval == 0 {
					log.Printf("loaded %v records from %v.wikipedia.org",
						i, req.wiki,
					)
				}
			}
		}

		log.Printf("finished loading %v records from %v.wikipedia.org", i, req.wiki)
	}
}

func findWikis(masks map[string]int64, value int64) []string {
	var wikis []string
	for wiki, mask := range masks {
		if mask&value == mask {
			wikis = append(wikis, wiki)
		}
	}

	return wikis
}

func getWord(t *algo.Trie, masks map[string]int64, w http.ResponseWriter, r *http.Request) {
	var node *algo.Trie

	response := wordResponse{}
	word := r.FormValue("word")

	response.Exists, node = t.Exists(word)
	if node != nil && node.Letter != 0 {
		rawCompletions := node.FindCompletions(word, maxCompletions)
		for _, rawCompletion := range rawCompletions {
			completion := completion{
				Word:  rawCompletion.Word,
				Wikis: findWikis(masks, rawCompletion.Node.Value),
			}
			response.Completions = append(response.Completions, completion)
		}
	}

	if node != nil {
		response.Wikis = findWikis(masks, node.Value)
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Println("can't marshal to json", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func main() {
	// Create our global trie
	trie := algo.NewTrie()

	// The list of wikis we want to download
	wikis := strings.Split(wikiCodes, "|")

	// Create a channel to concurrently download wikis
	dlChan := make(chan dlReq, len(wikis))

	// Map wiki to a bitmask
	wikiMasks := map[string]int64{}

	// Send the code and bitmask for each wiki to the downloader
	var mask int64 = 1
	for _, wiki := range wikis {
		dlChan <- dlReq{wiki: wiki, mask: mask}
		wikiMasks[wiki] = mask
		mask = mask << 1
	}

	// Create concurrent workers
	for i := 0; i < downloadWorkers; i++ {
		go download(dlChan, trie)
	}

	go loadStream(wikiMasks, trie)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/word", func(w http.ResponseWriter, r *http.Request) {
		getWord(trie, wikiMasks, w, r)
	})
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.Fatal(http.ListenAndServe(":53172", mux))
}
