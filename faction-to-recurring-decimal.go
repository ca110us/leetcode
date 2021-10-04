package main

import "strconv"

// 166 https://leetcode-cn.com/problems/fraction-to-recurring-decimal/
// 除法模拟
func fractionToDecimal(numerator int, denominator int) string {
	result := ""
	flag := "" // 符号位
	// 将符号位和数字分离开
	if denominator == 0 {
		return ""
	}
	if numerator == 0 {
		return "0"
	}
	if numerator*denominator > 0 {
		flag = ""
		if numerator < 0 {
			numerator = numerator * (-1)
			denominator = denominator * (-1)
		}
	} else {
		if numerator > 0 {
			flag = "-"
			denominator = denominator * (-1)
		} else {
			flag = "-"
			numerator = numerator * (-1)
		}
	}
	if (numerator % denominator) == 0 {
		return flag + strconv.Itoa(numerator/denominator)
	} else {
		// 定义整数和小数
		integer := 0
		// 若是被除数 > 除数，则算出商的整数部分
		for i := 1; numerator > i*denominator; i++ {
			integer++
		}
		result += strconv.Itoa(integer)
		result += "."
		numerator -= integer * denominator
		// 现在，被除数小于除数了，商的整数部分也算出来了，该算小数部分了
		hashMap := make(map[int]int) // 用于记录余数出现的位置，key是数字，value是其在quotient数组里的index
		var quotient []int           // 用于记录每次出现的小数部分商
		j := 0
		for numerator%denominator != 0 {
			if numerator < denominator {
				numerator *= 10
			}
			quotientRes := numerator / denominator
			numerator %= denominator
			if _, ok := hashMap[numerator]; !ok {
				// 表明没有循环，保存商和余数
				quotient = append(quotient, quotientRes)
				hashMap[numerator] = j
				j++
			} else {
				// 表明有循环
				if quotient[hashMap[numerator]] == quotientRes {
					for k := 0; k < hashMap[numerator] && k < len(quotient); k++ {
						result += strconv.Itoa(quotient[k])
					}
					result += "("
					for k := hashMap[numerator]; k < len(quotient); k++ {
						result += strconv.Itoa(quotient[k])
					}
					result += ")"
				} else {
					// 下面这段代码用于防止1/6这种情况的出现，即：余数重复出现，不一定代表就此循环
					for k := 0; k <= hashMap[numerator] && k < len(quotient); k++ {
						result += strconv.Itoa(quotient[k])
					}
					result += "("
					for k := hashMap[numerator] + 1; k < len(quotient); k++ {
						result += strconv.Itoa(quotient[k])
					}
					result += strconv.Itoa(quotientRes)
					result += ")"
				}
				break
			}
			// 表明除尽了
			if numerator == 0 {
				for _, item := range quotient {
					result += strconv.Itoa(item)
				}
				break
			}
		}
	}
	result = flag + result
	return result
}
