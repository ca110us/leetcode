package main

// 457 https://leetcode-cn.com/problems/circular-array-loop/submissions/
// 快慢指针相遇
func circularArrayLoop(nums []int) bool {
	ln := len(nums)
	for i := 0; i < ln; i++ {
		slow, fast := i, next(i, nums)
		// 判断同号和同向
		for nums[fast]*nums[i] > 0 && nums[next(fast, nums)]*nums[i] > 0 {
			if fast == slow {
				if slow == next(slow, nums) {
					break
				}
				return true
			}
			slow = next(slow, nums)
			fast = next(next(fast, nums), nums)
		}
	}
	return false
}

func next(cur int, nums []int) int {
	ln := len(nums)
	return ((cur+nums[cur])%ln + ln) % ln
}
