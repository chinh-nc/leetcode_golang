package jumpgame

func CanJump(nums []int) bool {
	step := 0
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if step < 0 {
			return false
		} else if nums[i] > step {
			step = nums[i]
		}
		step--
	}
	return true
}

/**
This is a Go code snippet that solves the "Jump Game II" problem from LeetCode. The problem is to determine the minimum number of jumps needed to reach the end of an array from the start.

Explanation: https://www.youtube.com/watch?v=dJ7sWiOoK7g
**/

func CanJump2(nums []int) int {
	count := 0
	var leftPoint, rightPoint int = 0, 0
	for rightPoint < len(nums)-1 {
		farthest := 0
		for i := leftPoint; i < rightPoint+1; i++ {
			farthest = max(farthest, i+nums[i])
		}
		leftPoint = rightPoint + 1
		rightPoint = farthest
		count++
	}
	return count
}