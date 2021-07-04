package main

import (
	"fmt"
	"sort"
	"strconv"
)

func countOfAtoms(formula string) string {
	strStack := []string{}
	numStack := []int{}

	str := ""
	for i := 0; i < len(formula); i++ {
		s := formula[i]
		// 数字
		num := 1
		if int(s) >= 49 && int(s) <= 57 {
			fmt.Println(string(s), num, str, "jjjjj")
			if str != "" {
				strStack = append(strStack, str)
				str = ""
			}
			num, _ = strconv.Atoi(string(s))
			numStack = append(numStack, num)
		}

		// 大写字母
		if int(s) >= 65 && int(s) <= 90 {
			if i == len(formula)-1 {
				str = string(s)
				strStack = append(strStack, str)
				numStack = append(numStack, num)
				continue
			}
			fmt.Println(string(s), num, str, "jjj000jj")
			if str != "" {
				strStack = append(strStack, str)
				str = ""
				numStack = append(numStack, num)
			}
			str = string(s)
		}

		// 小写字母
		if int(s) >= 97 && int(s) <= 122 {
			fmt.Println(str, string(s), "jjj99999")
			str = str + string(s)
		}

		// （
		if string(s) == "(" {
			if str != "" {
				strStack = append(strStack, str)
				str = ""
				numStack = append(numStack, num)
			}
			strStack = append(strStack, "(")
		}

		// ）
		if string(s) == ")" {
			if str != "" {
				strStack = append(strStack, str)
				str = ""
				numStack = append(numStack, num)
			}

			fmt.Println(strStack, numStack, "jsjdjdjj")
			ss := []string{}
			ns := []int{}

			base := 1
			// 下一个字符是数字
			if int(formula[i+1]) >= 49 && int(formula[i+1]) <= 57 {
				base, _ = strconv.Atoi(string(formula[i+1]))
				// 跳过下一个字符
				i = i + 1
			}

			for {
				// 直到取到 "("
				str := strStack[len(strStack)-1]
				strStack = strStack[:len(strStack)-1]
				if str == "(" {
					break
				}
				ss = append(ss, str)

				num := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				ns = append(ns, num)
			}

			fmt.Println(ns, ss, base, strStack, numStack, len(ns), "jdjdjdjdjjdijijie2222")
			ln := len(ns)
			for i := 0; i < ln; i++ {
				fmt.Println("kkkkkkk-------")
				// 归还新数字
				num := ns[len(ns)-1]
				ns = ns[:len(ns)-1]

				newNum := num * base
				numStack = append(numStack, newNum)

				// 归还字符串
				str := ss[len(ss)-1]
				ss = ss[:len(ss)-1]

				strStack = append(strStack, str)
			}

		}
	}

	m := map[string]int{}
	atoms := []string{}
	for i := 0; i < len(strStack); i++ {
		m[strStack[i]] = m[strStack[i]] + numStack[i]
		atoms = add(atoms, strStack[i])
	}

	sort.Strings(atoms)
	fmt.Println(m, atoms, "hhhuuuuuu")
	ans := ""
	for _, atom := range atoms {
		if m[atom] == 1 {
			ans += atom
			continue
		}
		ans += atom + strconv.Itoa(m[atom])
	}

	fmt.Println(ans)
	return ans
}

func add(s []string, e string) []string {
	for _, v := range s {
		if v == e {
			return s
		}
	}

	s = append(s, e)
	return s
}

func main() {
	countOfAtoms("Be32")
}
