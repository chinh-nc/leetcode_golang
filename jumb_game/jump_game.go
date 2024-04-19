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