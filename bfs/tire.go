package bfs

type node struct {
	isEnd    bool
	children map[byte]*node
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{root: &node{children: make(map[byte]*node)}}
}

// insert 插入, 已经存在返回false
func (t *trie) insert(word string) bool {
	n := t.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := n.children[c]; !ok {
			n.children[c] = &node{children: make(map[byte]*node)}
		}
		n = n.children[c]
	}
	if n.isEnd {
		return false
	}
	n.isEnd = true
	return true
}
