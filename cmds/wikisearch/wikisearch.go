package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/brnstz/algo"
)

const MAX_COMPLETIONS = 100
const WIKI_INDEX_URL = "https://dumps.wikimedia.org/enwiki/20170820/enwiki-20170820-pages-articles-multistream-index.txt.bz2"
const LOAD_LOG_INTERVAL = 1000000

type wordResponse struct {
	Exists      bool     `json:"exists"`
	Completions []string `json:"completions"`
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

	i := 0
	r := bufio.NewReader(os.Stdin)
	trie := algo.NewTrie(0)

	for ; ; i++ {
		text, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		text = strings.TrimSpace(text)

		strings := strings.SplitN(text, ":", 3)

		title := strings[2]

		trie.Add(title)

		if i > 0 && i%LOAD_LOG_INTERVAL == 0 {
			log.Printf("loaded %v lines", i)
		}

	}
	log.Printf("loading complete, %v lines\n", i)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/word", func(w http.ResponseWriter, r *http.Request) {
		getWord(trie, w, r)
	})
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.Fatal(http.ListenAndServe(":53172", mux))
}
