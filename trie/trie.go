package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

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
	// lastNode := t

	// For every letter in the word, ensure a trie
	// node exists.
	for _, letter := range word {

		child = node.EnsureChild(letter)

		// Run down the tree to the next node and add the next letter
		node = child
		// lastNode = node
	}

	// Setting as a leaf node indicates this is node is a word.
	node.Leaf = true
	// lastNode.Leaf = true
}

// Exists returns a boolean indicating whether this word exists or not in our
// trie.
func (t *Trie) Exists(word string) bool {
	node := t

	for _, letter := range word {

		// If at any point we don't have a child, then this word
		// doesn't exist
		node = node.FindChild(letter)
		if node == nil {
			return false
		}
	}

	// If the final letter isn't a leaf node, then it isn't a word
	return node.Leaf
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
	}
	log.Printf("loading complete, %v lines\n", i)
}
