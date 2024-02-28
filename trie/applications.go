package trie



func (t *Trie) PrefixTrie(s string) {
	for i := len(s) ; i > 0 ; i-- {
		t.Insert(s[:i])
	}
}

func (t *Trie) SuffixTrie(s string) {
	for i := 0 ; i < len(s) ; i++ {
		t.Insert(s[i:])
	}
}