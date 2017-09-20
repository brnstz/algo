package main

import (
	"bufio"
	"compress/bzip2"
	"encoding/json"
	"log"
	"net/http"

	"github.com/brnstz/algo"
)

const (
	MAX_COMPLETIONS   = 100
	LOAD_LOG_INTERVAL = 1000000
	DUMP_DATE         = "20178020"
	WIKI_INDEX_URL    = "http://dumps.wikimedia.your.org/%vwiki/%v/%vwiki-%v-pages-articles-multistream-index.txt.bz2"
	WIKIS             = "en|ceb|sv|de|nl|fr|ru|it|es|war|pl|vi|ja|pt|zh|uk|fa|ca|ar|no|sh|fi|hu|id|ko"
	DOWNLOAD_WORKERS  = 5
)

type wordResponse struct {
	Exists      bool     `json:"exists"`
	Completions []string `json:"completions"`
}

func download(url string, t *algo.Trie) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("can't download %v: %v", url, err)
		return
	}
	defer resp.Body.Close()

	reader, err := bzip2.NewReader(resp.Body)
	if err != nil {
		log.Println("can't read bzip2 data in %v: %v", url, err)
		return
	}

	s := bufio.NewScanner(reader)
	for s.Scan() {
		t.Add(s.Text())
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
	trie := algo.NewTrie(0)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/word", func(w http.ResponseWriter, r *http.Request) {
		getWord(trie, w, r)
	})
	mux.Handle("/", http.FileServer(http.Dir("static")))

	for i := 0; i < DOWNLOAD_WORKERS; i++ {
		go download()

	}

	log.Fatal(http.ListenAndServe(":53172", mux))
}
