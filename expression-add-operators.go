package main

import "strconv"

// 282 https://leetcode-cn.com/problems/expression-add-operators/
// dfs
func addOperators(num string, target int) []string {
	n := len(num)
	result := []string{}
	var dfs func(start int, res string, eval int, pre int)
	dfs = func(start int, res string, eval int, pre int) { //eval是上一层传下来表达式的真实计算值，pre 为上一层处理的数字连同符号
		if start == n {
			if eval == target {
				result = append(result, res)
			}
			return
		}
		for i := start; i < n; i++ { //bfs，每一个起点可以取一个或多个数
			if num[start] == '0' && i > start { // 多个数字情况下首数字为 0 要排除掉，011.... 这种情况
				return
			}
			number := num[start : i+1]         // 每层循环取长度为 start 到i的字符串作为本层新加入的操作数
			number2, _ := strconv.Atoi(number) //转化为数字
			if start == 0 {                    // 如果 start 为0，代表我们处理的是第一个数，那下一轮 dfs 不传入符号
				dfs(i+1, res+number, number2, number2)
			} else { // 已经是第二个操作数或者后面的了，必须给入符号到 res 中
				dfs(i+1, res+"+"+number, eval+number2, number2)
				dfs(i+1, res+"-"+number, eval-number2, -number2)
				dfs(i+1, res+"*"+number, eval-pre+pre*number2, pre*number2) // 乘法操作特殊，用题解中的公式方法
			}
		}
	}
	dfs(0, "", 0, 0)
	return result
}
