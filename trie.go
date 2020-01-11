package wordtrie

type Trie struct {
	Chr      rune
	IsWord   bool
	Parent   *Trie
	Children []*Trie
}

// GetChild returns the trie child at character chr
func (t *Trie) GetChild(chr rune) (*Trie, bool) {
	for _, node := range t.Children {
		if node.Chr == chr {
			return node, true
		}
	}

	return nil, false
}

// Insert inserts the word given into the trie
func (t *Trie) Insert(word string) {
	curTrie := t
	for i, r := range word {
		m, ok := curTrie.GetChild(r)
		if !ok {
			m = &Trie{r, i == len(word)-1, curTrie, []*Trie{}}
			curTrie.Children = append(curTrie.Children, m)
		}
		curTrie = m
	}
}

// Delete deletes a word from the trie.  It does not clean up the hierarchy.
func (t *Trie) Delete(word string) {
	endTrie, ok := t.TrieAt(word)
	if !ok {
		return
	}

	// set the trie to not a word
	endTrie.IsWord = false

	for {
		// if it has children or is a word, return
		if len(endTrie.Children) > 0 || endTrie.IsWord {
			return
		}

		endTrie = endTrie.Parent
		if endTrie == nil {
			return
		}

		newChildren := make([]*Trie, 0, len(endTrie.Children))
		for _, c := range endTrie.Children {
			if !c.IsWord && len(c.Children) == 0 {
				continue
			}

			newChildren = append(newChildren, c)
		}

		endTrie.Children = newChildren

	}
}

// TrieAt returns the trie object at the given word.
func (t *Trie) TrieAt(word string) (*Trie, bool) {
	curTrie := t
	for _, r := range word {
		m, ok := curTrie.GetChild(r)
		if !ok {
			return nil, false
		}
		curTrie = m
	}

	return curTrie, true
}

// BuildWord takes the current trie and moves upward via parents to retun the word
func (t *Trie) BuildWord() string {
	curTrie := t
	word := string(t.Chr)
	for curTrie.Parent != nil && curTrie.Parent.Chr != ' ' {
		word = string(curTrie.Parent.Chr) + word
		curTrie = curTrie.Parent
	}

	return word
}

// NewTrie initializes a new empty trie
func NewTrie() *Trie {
	return &Trie{' ', false, nil, []*Trie{}}
}
