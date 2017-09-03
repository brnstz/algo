package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const MAX_COMPLETIONS = 10

// Stolen from https://gist.github.com/moraes/2141121#gistcomment-1361598
type Queue []*Trie

func (q *Queue) Push(t *Trie) {
	*q = append(*q, t)
}

func (q *Queue) Pop() *Trie {
	t := (*q)[0]

	*q = (*q)[1:]

	return t
}

// Trie is a node in our Trie structure
type Trie struct {
	// The letter this Trie represents
	Letter rune

	// Is this node the end of a word?
	Leaf bool

	// A pointer to the next sibling of this node
	NextSibling *Trie

	// The head and tail of a linked list of children one level below
	// this node
	FirstChild *Trie
	LastChild  *Trie
}

// NewTrie creates a new trie node. Use 0 as the letter for the root node.
func NewTrie(letter rune) *Trie {
	t := &Trie{Letter: letter}

	return t
}

// FindChild finds a trie node for this rune at one level below t or returns
// nil
func (t *Trie) FindChild(letter rune) *Trie {
	child := t.FirstChild

	// Check all siblings for this letter
	for child != nil {
		if child.Letter == letter {
			return child
		}

		child = child.NextSibling
	}

	// Nope, couldn't find it
	return nil
}

// EnsureChild ensures that a trie node for this letter exists at one level
// below t
func (t *Trie) EnsureChild(letter rune) *Trie {

	// Can we find the child already?
	child := t.FindChild(letter)

	// If not, create a node and append it to the children list
	if child == nil {
		child = NewTrie(letter)

		if t.FirstChild == nil {

			// It's the first entry, we need to set head and tail
			t.FirstChild = child
			t.LastChild = child

		} else {

			// It's not the first entry
			t.LastChild.NextSibling = child
			t.LastChild = child
		}
	}

	return child
}

// Add a word to the trie
func (t *Trie) Add(word string) {
	var child *Trie

	// Start with our root trie
	node := t

	// For every letter in the word, ensure a trie
	// node exists.
	for _, letter := range word {

		child = node.EnsureChild(letter)

		// Run down the tree to the next node and add the next letter
		node = child
	}

	// Setting as a leaf node indicates this is node is a word.
	node.Leaf = true
}

// Exists returns a boolean indicating whether this word exists or not in our
// trie.
func (t *Trie) Exists(word string) (bool, *Trie) {
	node := t

	for _, letter := range word {

		// If at any point we don't have a child, then this word
		// doesn't exist
		node = node.FindChild(letter)
		if node == nil {
			return false, nil
		}
	}

	// If the final letter isn't a leaf node, then it isn't a word
	return node.Leaf, node
}

/*
func (t *Trie) findCompletionsAux(word string, max int) []string {
	node := t.FirstChild
	completions := []string{}

	// Do a breadth first search trying to find leaf nodes
	for child != nil {
		if child.Leaf {
			completions = append(completions, word+string(child.Letter))
		}

		if len(completions) >= max {
			return completions
		}
	}

	return completions
}
*/

func (t *Trie) FindCompletions(word string, max int) []string {
	var node *Trie
	var q Queue

	completions := []string{}
	node = t.FirstChild

	// Initialize q with root's children
	for node != nil {
		q.Push(node)

		node = node.NextSibling
	}

	for len(q) > 0 {
		node = q.Pop()

	}

	for node != nil {
		completions = node.findCompletions(word, max-len(completions))

		if len(completions) >= max {
			return completions
		}

	}

	return completions
}

type wordResponse struct {
	Exists      bool     `json:"exists"`
	Completions []string `json:"completions"`
}

func getWord(t *Trie, w http.ResponseWriter, r *http.Request) {
	var node *Trie

	response := wordResponse{}
	word := r.FormValue("word")

	response.Exists, node = t.Exists(word)
	response.Completions = node.FindCompletions(word, MAX_COMPLETIONS)

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
	trie := NewTrie(0)

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

		if i > 0 && i%1000000 == 0 {
			log.Printf("loaded %v lines", i)
		}

		if i >= 50000 {
			break
		}
	}
	log.Printf("loading complete, %v lines\n", i)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/word", func(w http.ResponseWriter, r *http.Request) {
		getWord(trie, w, r)
	})

	log.Fatal(http.ListenAndServe(":53172", mux))
}
