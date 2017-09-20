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
	maxCompletions  = 25
	dumpDate        = "20170820"
	wikiIndexURL    = "http://dumps.wikimedia.your.org/%vwiki/%v/%vwiki-%v-pages-articles-multistream-index.txt.bz2"
	wikiCodes       = "en|ceb|sv|de|nl|fr|ru|it|es|war|pl|vi|ja|pt|zh|uk|fa|ca|ar|no|sh|fi|hu|id|ko"
	downloadWorkers = 5
	titleField      = 3
)

var writeLock sync.Mutex

type wordResponse struct {
	Exists      bool     `json:"exists"`
	Completions []string `json:"completions"`
	Wikis       []string `json:"wikis"`
}

type dlReq struct {
	wiki string
	mask int64
}

func download(reqs chan dlReq, t *algo.Trie) {

	for req := range reqs {
		url := fmt.Sprintf(wikiIndexURL, req.wiki, dumpDate, req.wiki, dumpDate)

		log.Printf("begin downloading %v", url)

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("can't download %v: %v\n", url, err)
			return
		}
		defer resp.Body.Close()

		s := bufio.NewScanner(bzip2.NewReader(resp.Body))
		for s.Scan() {
			parts := strings.SplitN(s.Text(), ":", titleField)
			if len(parts) == titleField {
				writeLock.Lock()
				t.Add(parts[titleField-1], req.mask)
				writeLock.Unlock()
			}
		}
	}
}

func getWord(t *algo.Trie, masks map[string]int64, w http.ResponseWriter, r *http.Request) {
	var node *algo.Trie

	response := wordResponse{}
	word := r.FormValue("word")

	response.Exists, node = t.Exists(word)
	if node != nil && node.Letter != 0 {
		response.Completions = node.FindCompletions(word, maxCompletions)
	}

	if node != nil {
		for wiki, mask := range masks {
			if mask&node.Value == mask {
				response.Wikis = append(response.Wikis, wiki)
			}
		}
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

	mux := http.NewServeMux()
	mux.HandleFunc("/api/word", func(w http.ResponseWriter, r *http.Request) {
		getWord(trie, wikiMasks, w, r)
	})
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.Fatal(http.ListenAndServe(":53172", mux))
}
