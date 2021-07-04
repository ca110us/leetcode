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
	numStr := ""
	for i := 0; i < len(formula); i++ {
		s := formula[i]
		if int(s) >= 48 && int(s) <= 57 {
			numStr += string(s)
		} else {
			if string(s) != ")" {
				str += string(s)
			}
		}

		// // 大写字母
		// if int(s) >= 65 && int(s) <= 90 {
		// 	str += string(s)
		// }

		// // 小写字母
		// if int(s) >= 97 && int(s) <= 122 {
		// 	str += string(s)
		// }

		// // （
		// if string(s) == "(" {
		// 	str += string(s)
		// }

		// ）
		if string(s) == ")" {
			fmt.Println(strStack, numStack, "jsjdjdjj")
			ss := []string{}
			ns := []int{}

			base := 1
			// 下一个字符是数字
			if i+1 <= len(formula)-1 && int(formula[i+1]) >= 48 && int(formula[i+1]) <= 57 {
				j := 1
				numStr := ""
				for {
					if i+j <= len(formula)-1 && int(formula[i+j]) >= 48 && int(formula[i+j]) <= 57 {
						numStr += string(formula[i+j])
						j++
					} else {
						break
					}
				}
				base, _ = strconv.Atoi(numStr)
				// 跳过下一个字符
				i = i + (j - 1)
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

		if i == len(formula)-1 || !(int(formula[i+1]) >= 97 && int(formula[i+1]) <= 122) && !(int(formula[i+1]) >= 48 && int(formula[i+1]) <= 57) {
			if str == "" {
				continue
			}
			fmt.Println(str, numStr, "jskjdkjskdjfksjdkf000")
			strStack = append(strStack, str)
			num, _ := strconv.Atoi(numStr)
			if num == 0 {
				num = 1
			}
			if string(s) != "(" {
				numStack = append(numStack, num)
			}

			str = ""
			numStr = ""
		}
	}

	fmt.Println(strStack, numStack)
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
