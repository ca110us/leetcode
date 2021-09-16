package main

// 212 https://leetcode-cn.com/problems/word-search-ii/
// dfs 回溯
func findWords(board [][]byte, words []string) []string {
	trie := NewTrie()
	for _, v := range words {
		trie.Insert(&v, 0)
	}
	m, n := len(board), len(board[0])

	directions := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	res := make([]string, 0)
	var backtrack func(i, j, curIdx int, last *Trie)
	backtrack = func(i, j, curIdx int, last *Trie) {
		curTrie := last.Children[curIdx]
		if curTrie.Word != "" {
			res = append(res, curTrie.Word)
			curTrie.Word = ""
		}
		letter := board[i][j]
		board[i][j] = '#'
		for _, v := range directions {
			newI, newJ := i+v[0], j+v[1]
			if newI >= 0 && newI < m && newJ >= 0 && newJ < n && board[newI][newJ] != '#' {
				if curTrie.Children[int(board[newI][newJ]-'a')] != nil {
					backtrack(newI, newJ, int(board[newI][newJ]-'a'), curTrie)
				}
			}
		}
		board[i][j] = letter
		if curTrie.Count == 0 {
			last.Children[curIdx] = nil
		}
	}

	for i := range board {
		for j := range board[i] {
			idx := int(board[i][j] - 'a')
			if trie.Children[idx] != nil {
				backtrack(i, j, idx, trie)
			}
		}
	}
	return res
}

type Trie struct {
	Word     string
	Children [26]*Trie
	Count    int
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word *string, idx int) {
	if len(*word) == idx {
		t.Word = *word
		return
	}

	if t.Children[int((*word)[idx]-'a')] == nil {
		t.Children[int((*word)[idx]-'a')] = new(Trie)
		t.Count++
	}

	t.Children[int((*word)[idx]-'a')].Insert(word, idx+1)
}
