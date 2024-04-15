package removeduplicate

// RemoveDuplicates removes duplicates in-place using pointer.
// item in nums is a pointer to an int of original nums

func RemoveDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
