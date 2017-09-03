package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

var letters = 0
var trieCalls = 0
var maxChildren = 0
var maxChildrenNode *Trie

type Trie struct {
	Letter   rune
	Children map[rune]*Trie
	Leaf     bool
}

// NewTrie creates a new trie node. Use 0 as the letter for the root node.
func NewTrie(letter rune) *Trie {
	t := &Trie{}
	t.Letter = letter

	trieCalls += 1

	return t
}

func (t *Trie) EnsureChild(runeValue rune) *Trie {
	var (
		child  *Trie
		exists bool
	)

	if t.Children == nil {
		// If there's no map, then there's definitely no child
		t.Children = map[rune]*Trie{}
		child = NewTrie(runeValue)
		t.Children[runeValue] = child

	} else {

		// If the child exists, use it, otherwise create it
		child, exists = t.Children[runeValue]
		if !exists {
			child = NewTrie(runeValue)
			t.Children[runeValue] = child
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
		letters += 1

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
	var exists bool
	node := t

	for _, runeValue := range word {
		node, exists = node.Children[runeValue]

		// If at any point we don't have a child, then this word
		// doesn't exist
		if !exists {
			return false
		}
	}

	// If the final letter isn't a leaf node, then it isn't a word
	return node.Leaf
}

func (t *Trie) countNodes() (int, int) {
	var fullNodes, emptyNodes int

	if t.Children != nil {
		fullNodes = 1
	} else {
		emptyNodes = 1
	}

	if len(t.Children) > maxChildren && t.Letter != 0 {
		maxChildren = len(t.Children)
		maxChildrenNode = t
	}

	for _, v := range t.Children {
		nowFull, nowEmpty := v.countNodes()

		fullNodes += nowFull
		emptyNodes += nowEmpty
	}

	return fullNodes, emptyNodes
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

		if i > 5000000 {
			break
		}
	}

	full, empty := trie.countNodes()

	fmt.Printf("letters: %v\n", letters)
	fmt.Printf("full nodes: %v\n", full)
	fmt.Printf("empty nodes: %v\n", empty)
	fmt.Printf("trie calls: %v\n", trieCalls)
	fmt.Printf("max children: %v\n", maxChildren)
	fmt.Printf("max children nodes: %v\n", maxChildrenNode)

	pprof.WriteHeapProfile(f)
	f.Close()
}
