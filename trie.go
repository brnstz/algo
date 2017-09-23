package algo

// Trie is a node in our Trie structure
type Trie struct {
	// The letter this Trie represents
	Letter rune

	// Value is a bitmask that determines what we can do
	// with this node. A zero value must represent a non-word, whereas
	// a non-zero can be interpreted by the client however it wants.
	Value int64

	// A pointer to the next sibling of this node
	Sibling *Trie

	// A pointer to the first child of this node
	Child *Trie
}

// NewTrie creates a new trie root with 0 as the value.
func NewTrie() *Trie {
	return &Trie{Letter: 0}
}

// newTrieNode creates a new trie node.
func newTrieNode(letter rune) *Trie {
	return &Trie{Letter: letter}
}

// findChild finds a trie node for this rune at one level below t or returns
// nil
func (t *Trie) findChild(letter rune) *Trie {
	child := t.Child

	// Check all siblings for this letter
	for child != nil {
		if child.Letter == letter {
			return child
		}

		child = child.Sibling
	}

	// Nope, couldn't find it
	return nil
}

// ensureChild ensures that a trie node for this letter exists at one level
// below t. Returns the node itself, whether this node was newly created,
// and how many siblings the node has.
func (t *Trie) ensureChild(letter rune) (*Trie, bool, int) {
	var (
		siblings int
		newNode  bool
	)

	// Can we find the child already?
	child := t.findChild(letter)

	// If not, create a node and append it to the children list
	if child == nil {
		child = newTrieNode(letter)
		newNode = true

		if t.Child == nil {
			// If it's the first child, just set it
			t.Child = child

		} else {
			// Otherwise, find the tail and set its
			// sibling to the new child
			tail := t.Child

			for tail.Sibling != nil {
				tail = tail.Sibling
				siblings++
			}

			tail.Sibling = child
		}
	}

	return child, newNode, siblings
}

// Add a word to the trie. Returns how many new nodes were created, and the
// maximum number of siblings a node has.
func (t *Trie) Add(word string, value int64) (int, int) {
	var (
		child                           *Trie
		newNode                         bool
		newNodes, maxSiblings, siblings int
	)

	// Start with our root trie
	node := t

	// For every letter in the word, ensure a trie
	// node exists.
	for _, letter := range word {

		// Create new child
		child, newNode, siblings = node.ensureChild(letter)

		// If we created a new node, record that
		if newNode {
			newNodes++
		}

		// If we have a new max siblings count, record that
		if siblings > maxSiblings {
			maxSiblings = siblings
		}

		// Run down the tree to the next node and add the next letter
		node = child
	}

	// Set new value of this node by running OR on existing value
	node.Value = node.Value | value

	return newNodes, maxSiblings
}

// Exists returns a boolean indicating whether this word exists or not in our
// trie. It also returns the Trie node when found.
func (t *Trie) Exists(word string) (bool, *Trie) {
	node := t

	for _, letter := range word {

		// If at any point we don't have a child, then this word
		// doesn't exist
		node = node.findChild(letter)
		if node == nil {
			return false, nil
		}
	}

	// If the final letter isn't a leaf node, then it isn't a word
	return node.Value > 0, node
}

// trieWord is a Trie node and the word up until that node. Eg, if we were
// storing "goodbye", the word might be "goodb" and the trie node might be
// the letter "y". This allows us to use a queue to run a breadth first
// search in FindCompletions
type trieWord struct {
	word string
	trie *Trie
}

// Stolen from https://gist.github.com/moraes/2141121#gistcomment-1361598
type queue []trieWord

func (q *queue) Push(t trieWord) {
	*q = append(*q, t)
}

func (q *queue) Pop() trieWord {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}

type Completion struct {
	Word string
	Node *Trie
}

// FindCompletions does a breadth-first search below this trie node, and
// finds up to max completed words under it.
func (t *Trie) FindCompletions(word string, max int) []Completion {
	var (
		child       *Trie
		tw          trieWord
		q           queue
		completions []Completion
	)

	// Initialize q with ourselves
	q.Push(trieWord{word: word, trie: t})

	// While we still have nodes in our queue
	for len(q) > 0 {

		// Get the word and trie node off the queue
		tw = q.Pop()

		// Check for children that complete a word
		child = tw.trie.Child
		for child != nil {
			childWord := tw.word + string(child.Letter)

			// If it's a word, add it to our words
			if child.Value > 0 {
				completion := Completion{
					Word: childWord,
					Node: child,
				}
				completions = append(completions, completion)
			}

			// If we have enough words, then stop
			if len(completions) >= max {
				return completions
			}

			// Add child to queue to process its children
			q.Push(trieWord{word: childWord, trie: child})

			// Try the next sibling
			child = child.Sibling
		}
	}

	return completions
}
