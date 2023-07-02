package models

type TrieNode struct {
	children    [26]*TrieNode
	frequency   int
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

// Inc increments the frequency of the word.
func (t *Trie) Inc(word []byte) {
	node := t.root
	for _, letter := range word {
		idx := letter - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.frequency++
	node.isEndOfWord = true
}

func (t *Trie) ToSlice() []Word {
	ret := make([]Word, 0)
	dfs(t.root, []byte{}, &ret)
	return ret
}

func dfs(node *TrieNode, currWord []byte, result *[]Word) {
	if node == nil {
		return
	}

	if node.isEndOfWord {
		w := Word{
			Bytes:     make([]byte, len(currWord)),
			Frequency: node.frequency,
		}
		copy(w.Bytes, currWord)
		*result = append(*result, w)
	}

	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			letter := byte(i) + 'a'
			dfs(node.children[i], append(currWord, letter), result)
		}
	}
}
