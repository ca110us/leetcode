package main

import (
	"sort"
	"strconv"
)

// 726 https://leetcode-cn.com/problems/number-of-atoms/
// 这题可以参照 Dijkstra 的双栈算术表达式求值算法（《算法-第四版》 80 页），其实不难，但是很烦
func isNum(s byte) bool         { return s >= '0' && s <= '9' }
func isLowerLatter(s byte) bool { return s >= 'a' && s <= 'z' }
func isUpperLatter(s byte) bool { return s >= 'A' && s <= 'Z' }

func countOfAtoms(formula string) string {
	strStack := []string{}
	numStack := []int{}

	var pre byte
	for i := 0; i < len(formula); i++ {
		s := formula[i]
		haveNext := i+1 < len(formula)

		// 大写字母处理
		if isUpperLatter(s) {
			strStack = append(strStack, string(s))
			if (haveNext && !isNum(formula[i+1]) && !isLowerLatter(formula[i+1])) || !haveNext {
				numStack = append(numStack, 1)
			}
			pre = s
			continue
		}

		// 小写字母处理
		if isLowerLatter(s) {
			strStack[len(strStack)-1] = strStack[len(strStack)-1] + string(s)
			if (haveNext && !isNum(formula[i+1]) && !isLowerLatter(formula[i+1])) || !haveNext {
				numStack = append(numStack, 1)
			}
			pre = s
			continue
		}

		// 数字处理
		if isNum(s) {
			if isNum(pre) {
				numStr := strconv.Itoa(numStack[len(numStack)-1])
				numStr += string(s)

				newNum, _ := strconv.Atoi(numStr)
				numStack[len(numStack)-1] = newNum
			} else {
				num, _ := strconv.Atoi(string(s))
				numStack = append(numStack, num)
			}
			pre = s
			continue
		}

		// 右括号处理 取出 strStack 中的所有原子，直到遇到 "(" ，从 numStack 中取出相同数量元素分别和 ")" 后面数字相乘
		// 再将新的 str 和 num 压回 strStack 和 numStack
		if string(s) == ")" {
			ss := []string{}
			ns := []int{}

			base := 1
			// 取出右括号后所有的数字
			if i+1 < len(formula) && isNum(formula[i+1]) {
				j := 1
				numStr := ""
				for i+1 < len(formula) && isNum(formula[i+1]) {
					if i+j <= len(formula)-1 && isNum(formula[i+j]) {
						numStr += string(formula[i+j])
						j++
					} else {
						break
					}
				}
				base, _ = strconv.Atoi(numStr)
				// 跳过 j-1 个字符
				i = i + (j - 1)
			}

			for {
				str := strStack[len(strStack)-1]
				strStack = strStack[:len(strStack)-1]
				// 直到取到 "("
				if str == "(" {
					break
				}
				ss = append(ss, str)

				num := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				ns = append(ns, num)
			}

			ln := len(ns)
			for i := 0; i < ln; i++ {
				// 压入新数字
				num := ns[len(ns)-1]
				ns = ns[:len(ns)-1]

				newNum := num * base
				numStack = append(numStack, newNum)

				// 压入字符
				str := ss[len(ss)-1]
				ss = ss[:len(ss)-1]

				strStack = append(strStack, str)
			}
			pre = s
			continue
		}

		strStack = append(strStack, string(s))
		pre = s
	}

	m := map[string]int{}
	atoms := []string{}
	for i := 0; i < len(strStack); i++ {
		m[strStack[i]] = m[strStack[i]] + numStack[i]
		atoms = add(atoms, strStack[i])
	}

	sort.Strings(atoms)

	ans := ""
	for _, atom := range atoms {
		if m[atom] == 1 {
			ans += atom
			continue
		}
		ans += atom + strconv.Itoa(m[atom])
	}

	return ans
}

// 去重
func add(s []string, e string) []string {
	for _, v := range s {
		if v == e {
			return s
		}
	}

	s = append(s, e)
	return s
}
