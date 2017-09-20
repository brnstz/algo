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
	MAX_COMPLETIONS  = 25
	DUMP_DATE        = "20170820"
	WIKI_INDEX_URL   = "http://dumps.wikimedia.your.org/%vwiki/%v/%vwiki-%v-pages-articles-multistream-index.txt.bz2"
	WIKIS            = "en|ceb|sv|de|nl|fr|ru|it|es|war|pl|vi|ja|pt|zh|uk|fa|ca|ar|no|sh|fi|hu|id|ko"
	DOWNLOAD_WORKERS = 5
)

var writeLock sync.Mutex

type wordResponse struct {
	Exists      bool     `json:"exists"`
	Completions []string `json:"completions"`
}

func download(urls chan string, t *algo.Trie) {

	for url := range urls {
		log.Printf("begin downloading %v", url)

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("can't download %v: %v\n", url, err)
			return
		}
		defer resp.Body.Close()

		s := bufio.NewScanner(bzip2.NewReader(resp.Body))
		for s.Scan() {
			parts := strings.SplitN(s.Text(), ":", 3)
			if len(parts) > 2 {
				writeLock.Lock()
				t.Add(parts[2])
				writeLock.Unlock()
			}
		}
		log.Printf("finished downloading %v", url)
	}
}

func getWord(t *algo.Trie, w http.ResponseWriter, r *http.Request) {
	var node *algo.Trie

	response := wordResponse{}
	word := r.FormValue("word")

	response.Exists, node = t.Exists(word)
	if node != nil && node.Letter != 0 {
		response.Completions = node.FindCompletions(word, MAX_COMPLETIONS)
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
	wikis := strings.Split(WIKIS, "|")

	// Create a channel to concurrently download wikis
	dlchan := make(chan string, len(wikis))

	// Send the URL for each wiki to the downloader
	for _, wiki := range wikis {
		dlchan <- fmt.Sprintf(WIKI_INDEX_URL, wiki, DUMP_DATE, wiki, DUMP_DATE)
	}

	// Create concurrent workers
	for i := 0; i < DOWNLOAD_WORKERS; i++ {
		go download(dlchan, trie)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/word", func(w http.ResponseWriter, r *http.Request) {
		getWord(trie, w, r)
	})
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.Fatal(http.ListenAndServe(":53172", mux))
}
