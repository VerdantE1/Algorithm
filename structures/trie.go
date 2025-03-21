package structures

// TrieNode 定义前缀树的节点结构
type TrieNode struct {
	isEnd    bool
	children map[rune]*TrieNode
}

// Trie 定义前缀树
type Trie struct {
	root *TrieNode
}

// NewTrie 创建一个新的前缀树
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			isEnd:    true,
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert 插入一个字符串到前缀树
func (t *Trie) Insert(word string) {
	root := t.root
	for _, ch := range word {
		if root.children[ch] == nil { //若不存在该子节点
			root.children[ch] = &TrieNode{
				isEnd:    false,
				children: make(map[rune]*TrieNode),
			}
		}
		root = root.children[ch] //若存在则递归
	}
	root.isEnd = true
}

// Search 从前缀树中找一个字符串
func (t *Trie) Search(word string) bool {
	root := t.root
	for _, ch := range word {
		if root.children[ch] == nil {
			return false
		}
		if root.children[ch].isEnd == true {
			return true
		}
		root = root.children[ch]
	}
	//if root.isEnd == false return false
	//return true
	return root.isEnd
}

// StartWith 从前缀树中找到一个前缀
func (t *Trie) StartsWith(prefix string) bool {
	root := t.root
	for _, ch := range prefix {
		if root.children[ch] == nil {
			return false
		}
		root = root.children[ch]
	}
	return true
}
