package removetriplicate

// WAY 1
// func RemoveTriplicate(nums []int) int {
// 	currentIndex := 1
// 	canInsert := 1
// 	for i := 1; i < len(nums); i++ {
// 		isDuplicated := nums[i] == nums[i-1]
// 		if !isDuplicated || canInsert < 3 {
// 			nums[currentIndex] = nums[i]
// 			currentIndex++
// 		}
// 		if isDuplicated {
// 			canInsert++
// 		} else {
// 			canInsert = 1
// 		}
// 	}

// 	return currentIndex

// }
func RemoveTriplicate(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	currentIndex := 2
	for i := 2; i < l; i++ {
		nums[currentIndex] = nums[i]
		if nums[i] != nums[currentIndex - 2] {
			currentIndex++
		}
	}

	return currentIndex

}