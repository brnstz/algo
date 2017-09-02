package main

import "fmt"

type Trie struct {
	Letter   rune
	Children map[rune]*Trie
	Leaf     bool
}

// NewTrie creates a new trie node. Use 0 as the letter for the root node.
func NewTrie(letter rune) *Trie {
	t := &Trie{}
	t.Letter = letter
	t.Children = map[rune]*Trie{}

	return t
}

// Add a word to the trie
func (t *Trie) Add(word string) {
	// Start with our root trie
	node := t

	// For every letter in the word, ensure a trie
	// node exists.
	for _, runeValue := range word {

		// If the child exists, use it, otherwise create it
		child, exists := node.Children[runeValue]
		if !exists {
			child = NewTrie(runeValue)
			node.Children[runeValue] = child
		}

		// Run down the tree to the next node and add the next letter
		node = child
	}

	// Setting as a leaf node indicates this is node is a word.
	node.Leaf = true
	fmt.Printf("%c %v\n", node.Letter, node.Leaf)
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

func main() {
	//r := bufio.NewReader(os.Stdin)

	trie := NewTrie(0)

	trie.Add("hello")
	trie.Add("goodbye")
	fmt.Println(trie.Exists("blah"))
	fmt.Println(trie.Exists("hello"))

	/*
		for {
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
			fmt.Println(trie)
		}
	*/
}
