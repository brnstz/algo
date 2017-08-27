package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Trie struct {
	Letter   rune
	Children map[rune]Trie
	Leaf     bool
}

func NewTrie(letter rune) Trie {
	t := Trie{}
	t.Letter = letter
	t.Children = map[rune]Trie{}

	return t
}

func (t Trie) Add(word string) {
	node := t
	lastNode := t

	for _, runeValue := range word {
		lastNode = node

		child, exists := node.Children[runeValue]
		if !exists {
			child = NewTrie(runeValue)
			node.Children[runeValue] = child
		}

		node = child
	}

	lastNode.Leaf = true
}

func main() {
	r := bufio.NewReader(os.Stdin)

	trie := NewTrie(0)

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

}
