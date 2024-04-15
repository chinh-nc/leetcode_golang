package rotate

import "fmt"

func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	fmt.Println(k)
	reverse(nums, 0, n-k-1)
	fmt.Println(nums)
	reverse(nums, n-k, n-1)
	fmt.Println(nums)
	reverse(nums, 0, n-1)
	fmt.Println(nums)
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, start int, end int) {
	for start < end {
		swap(nums, start, end)
		start++
		end--
	}
}
