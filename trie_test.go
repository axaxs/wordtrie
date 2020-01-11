package wordtrie

import (
	"testing"
)

func TestNewTrie(t *testing.T) {
	n := NewTrie()
	if n == nil {
		t.Fatal("New trie is nil!")
	}
}

func TestGetChild(t *testing.T) {
	n := NewTrie()
	n.Insert("hi")
	_, ok := n.GetChild('h')
	if !ok {
		t.Fatal("Trying to get child that should exist failed")
	}
}

func TestGetTrieAt(t *testing.T) {
	n := NewTrie()
	n.Insert("hi")
	c, ok := n.TrieAt("hi")
	if !ok {
		t.Fatal("Trying to get trie that should exist failed")
	}
	if c.Parent == nil || c.Chr != 'i' || !c.IsWord {
		t.Fatal("Trie should have parent, Chr 'i', and be marked IsWord, but is not")
	}
}

func TestDelete(t *testing.T) {
	n := NewTrie()
	n.Insert("hi")
	n.Delete("hi")
	_, ok := n.TrieAt("hi")
	if ok {
		t.Fatal("Trying to get trie that should not exist succeeded")
	}

	_, ok = n.TrieAt("h")
	if ok {
		t.Fatal("Trying to get parent trie that should not exist succeeded")
	}
}

func TestBuildWord(t *testing.T) {
	n := NewTrie()
	n.Insert("hi")
	c, ok := n.TrieAt("hi")
	if !ok {
		t.Fatal("Trying to get trie that should exist failed")
	}
	x := c.BuildWord()
	if x != "hi" {
		t.Fatalf("Expecting 'hi', got '%s'", x)
	}
}
