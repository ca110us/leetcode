package main

import "math"

// 909 https://leetcode-cn.com/problems/snakes-and-ladders
// BFS 可以转变成一维数组处理，也可以根据 index 转化成二维数组中的对应 (getValue) 值
func snakesAndLadders(board [][]int) int {
	var realBoard []int
	realBoard = append(realBoard, 0)

	left := true
	for i := len(board) - 1; i >= 0; i-- {
		v := board[i]
		if !left {
			realBoard = append(realBoard, reverse(v)...)
		} else {
			realBoard = append(realBoard, v...)
		}
		left = !left
	}

	q := []int{1}
	max := len(board) * len(board)
	map2step := make(map[int]int, max+1)
	map2step[1] = 0

	for len(q) > 0 {
		now := q[0]
		q = q[1:]
		step := map2step[now]

		for i := 1; i <= 6; i++ {
			nxt := now + i
			if nxt > max {
				break
			}

			if realBoard[nxt] != -1 {
				nxt = realBoard[nxt]
			}

			if nxt == max {
				return step + 1
			}

			if _, ok := map2step[nxt]; !ok {
				map2step[nxt] = step + 1
				q = append(q, nxt)
			}
		}
	}

	return -1
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func getValue(board [][]int, index int) int {
	n := len(board)
	line := n - int(math.Ceil(float64(index)/float64(n)))
	col := 0
	if (n-1-line)&1 == 0 {
		col = index%n - 1
		if index%n == 0 {
			col = n - 1
		}
	} else {
		col = n - (index % n)
		if index%n == 0 {
			col = 0
		}
	}
	return board[line][col]
}
