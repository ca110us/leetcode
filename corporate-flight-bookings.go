package main

// 1109 https://leetcode-cn.com/problems/corporate-flight-bookings/submissions/
// 不明所以 按照题目随手一打 就过了 ...
func corpFlightBookings(bookings [][]int, n int) []int {
	seats := make([]int, n)

	for _, booking := range bookings {
		start, end := booking[0]-1, booking[1]-1
		for i := start; i <= end; i++ {
			seats[i] += booking[2]
		}
	}

	return seats
}
