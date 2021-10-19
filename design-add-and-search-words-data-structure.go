package main

// 211 https://leetcode-cn.com/problems/design-add-and-search-words-data-structure/
// 字典树
type WordDictionary struct {
	arr  [26]*WordDictionary
	isOk bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	tmp := this
	for _, v := range word {
		v -= 'a'
		if tmp.arr[v] == nil {
			tmp.arr[v] = &WordDictionary{}
		}
		tmp = tmp.arr[v]
	}
	tmp.isOk = true
}

func dfs(node *WordDictionary, word string, index int) bool {
	if node == nil {
		return false
	}
	if len(word) == index {
		if node.isOk {
			return true
		}
		return false
	}
	ch := word[index]
	if ch != '.' {
		ch -= 'a'
		node = node.arr[ch]
		return dfs(node, word, index+1)
	} else {
		res := false
		for i := 0; i < 26; i++ {
			if node.arr[i] != nil {
				res = res || dfs(node.arr[i], word, index+1)
			}
		}
		return res
	}
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this, word, 0)
}
