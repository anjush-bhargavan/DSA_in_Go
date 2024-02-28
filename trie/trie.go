package trie

type TrieNode struct {
	Children map[byte]*TrieNode
	IsEnd    bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() Trie {
	return Trie{Root: &TrieNode{Children: make(map[byte]*TrieNode)}}
}


func (t *Trie) Insert(s string) {
	curr := t.Root
	for i := 0 ; i < len(s) ; i++ {
		if curr.Children[s[i]] == nil  {
			curr.Children[s[i]] = &TrieNode{Children: map[byte]*TrieNode{}}
		}
		curr = curr.Children[s[i]]
	}
	curr.IsEnd = true
}

func (t *Trie) Contains(s string) bool {
	curr := t.Root

	for i := 0; i < len(s); i++ {
		if curr.Children[s[i]] == nil {
			return false
		}
		curr = curr.Children[s[i]]
	}
	return curr.IsEnd
} 


func (t *Trie) Remove(s string) {
	t.deleteRecursive(t.Root,s,0)
}

func (t *Trie) deleteRecursive(node *TrieNode,s string,i int) *TrieNode {
	if node == nil {
		return nil 
	}

	if i == len(s) {
		if node.IsEnd {
			node.IsEnd = false
		}


		if len(node.Children) == 0 {
			return nil 
		}

		return node
	}

	c := s[i]
	node.Children[c] = t.deleteRecursive(node.Children[c],s,i+1)

	if len(node.Children) == 0 && !node.IsEnd {
		return nil 
	}
	
	return node
}