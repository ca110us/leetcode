package main

import "strings"

// 68 https://leetcode-cn.com/problems/text-justification/
// 模拟
func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	size, cnt := 0, 0

	getRowWord := func(cnt, end, blankCnt int) string {
		var builder strings.Builder
		avgBlankCnt, leftBlankCnt := blankCnt, 0

		if cnt > 1 {
			avgBlankCnt = blankCnt / (cnt - 1)
			leftBlankCnt = blankCnt % (cnt - 1)
		}
		var blankBuilder strings.Builder
		for i := 0; i < avgBlankCnt; i++ {
			blankBuilder.WriteByte(' ')
		}

		for i := end - cnt + 1; i <= end; i++ {
			builder.WriteString(words[i])
			if i < end {
				builder.WriteString(blankBuilder.String())
				if leftBlankCnt > 0 {
					builder.WriteByte(' ')
					leftBlankCnt--
				}
			}
		}
		if cnt == 1 {
			builder.WriteString(blankBuilder.String())
		}
		return builder.String()
	}
	idx := 0
	for idx < len(words) {
		word := words[idx]
		if size+len(word) <= maxWidth-cnt {
			size += len(word)
			cnt++
			idx++
		} else {
			rowWord := getRowWord(cnt, idx-1, maxWidth-size)
			result = append(result, rowWord)
			size = 0
			cnt = 0
		}
	}
	var lastRow strings.Builder
	// 处理最后一行
	if size > 0 {
		idx = len(words) - 1

		for i := idx - cnt + 1; i <= idx; i++ {
			lastRow.WriteString(words[i])
			if i < idx {
				// 最后一行一个空格
				lastRow.WriteByte(' ')
			}
		}
	}
	leftCnt := maxWidth - lastRow.Len()
	for i := 0; i < leftCnt; i++ {
		lastRow.WriteByte(' ')
	}
	result = append(result, lastRow.String())
	return result
}
