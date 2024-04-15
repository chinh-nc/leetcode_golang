package removeelement

import "fmt"

// RemoveElement removes all occurrences of val in nums in-place.
// It returns the length of nums after removal.
// using pointer
// item in nums is a pointer to an int of original nums

func RemoveElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	fmt.Println(&nums)
	return count
}