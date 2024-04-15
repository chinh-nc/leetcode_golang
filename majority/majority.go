package majority

import "sort"

func Majority(nums []int) int {
	majorityNum := int(len(nums) / 2)
	sort.Ints(nums)
	return nums[majorityNum]
}