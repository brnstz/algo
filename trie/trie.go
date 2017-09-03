package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

type Trie struct {
	Letter rune
	Leaf   bool

	NextSibling *Trie

	FirstChild *Trie
	LastChild  *Trie
}

// NewTrie creates a new trie node. Use 0 as the letter for the root node.
func NewTrie(letter rune) *Trie {
	t := &Trie{Letter: letter}

	return t
}

func (t *Trie) FindChild(runeValue rune) *Trie {
	child := t.FirstChild

	for child != nil {
		if child.Letter == runeValue {
			return child
		}

		child = child.NextSibling
	}

	return nil
}

func (t *Trie) EnsureChild(runeValue rune) *Trie {
	child := t.FindChild(runeValue)

	if child == nil {
		child = NewTrie(runeValue)

		// Two possibilities

		if t.FirstChild == nil {
			// It's the first entry
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
	lastNode := t

	// For every letter in the word, ensure a trie
	// node exists.
	for _, runeValue := range word {

		child = node.EnsureChild(runeValue)

		// Run down the tree to the next node and add the next letter
		node = child
		lastNode = node
	}

	// Setting as a leaf node indicates this is node is a word.
	lastNode.Leaf = true
}

// Does this word exist?
func (t *Trie) Exists(word string) bool {
	node := t

	for _, runeValue := range word {
		node = node.FindChild(runeValue)

		// If at any point we don't have a child, then this word
		// doesn't exist
		if node == nil {
			return false
		}
	}

	// If the final letter isn't a leaf node, then it isn't a word
	return node.Leaf
}

func main() {
	r := bufio.NewReader(os.Stdin)
	trie := NewTrie(0)

	/*
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	*/

	f, err := os.Create("memprofile.out")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; ; i++ {
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

		if i%1000000 == 0 {
			log.Println(i)
			t1 := time.Now()
			found := trie.Exists("Americans with Disabilities Act of 1990/Findings and Purposes")
			t2 := time.Now()
			log.Printf("%v %v\n", found, t2.Sub(t1))
		}

		/*
			if i > 5000000 {
				break
			}
		*/
	}

	pprof.WriteHeapProfile(f)
	f.Close()
}
