package structures

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	// 测试插入和搜索
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Expected true, got false for Search(\"apple\")")
	}
	if trie.Search("app") {
		t.Errorf("Expected false, got true for Search(\"app\")")
	}

	// 测试前缀
	if !trie.StartsWith("app") {
		t.Errorf("Expected true, got false for StartsWith(\"app\")")
	}

	// 插入更多字符串
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Expected true, got false for Search(\"app\")")
	}

	// 测试删除
	trie.Delete("app")
	if trie.Search("app") {
		t.Errorf("Expected false, got true for Search(\"app\") after deletion")
	}
	if !trie.Search("apple") {
		t.Errorf("Expected true, got false for Search(\"apple\") after deletion of \"app\"")
	}
}
