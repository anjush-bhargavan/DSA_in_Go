package trie

type TrieNode struct {
	children 	map[byte]*TrieNode
	isEnd 		bool
}

type Trie struct {
	root *TrieNode
}

func NewSuffixTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[byte]*TrieNode)}}
}

func (t *Trie) Trie(str string) {
	t.popuplateSuffixTrie(str)
}

func (t *Trie) popuplateSuffixTrie(str string) {
	for i := 0 ; i < len(str) ; i++ {
		t.insertSubstringStartingAt(i,str)
	}
}

func (t *Trie) insertSubstringStartingAt(index int, str string) {
	node := t.root
	for i := index ; i<len(str) ; i++ {
		if node.children[str[i]] == nil {
			node.children[str[i]] = &TrieNode{children: map[byte]*TrieNode{}}
		}
		node = node.children[str[i]]
	}
	node.isEnd = true
}

func (t *Trie) Contains(str string) bool {
	node := t.root
	for i := 0 ; i < len(str) ; i++ {
		if node.children[str[i]]== nil {
			return false
		}
		node = node.children[str[i]]
	}
	
	return node.isEnd
}


/*0000000000000000000000000000000000000000000000000000000000000000000000000000000000*/

func (t *Trie) PrefixString(str string) {
	node := t.root
	for i := 0 ; i<len(str) ; i++ {
		if node.children[str[i]] == nil {
			// newNode := &TrieNode{}
			node.children[str[i]] = &TrieNode{children: map[byte]*TrieNode{}}
		}
		node = node.children[str[i]]
	}
	node.isEnd = true
}