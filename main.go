package main

import (
	"fmt"
	"leetcode/rotate"
)

func main() {
	nums := []int{1,2,3,4,5,6,7}
	fmt.Println(nums)
	rotate.Rotate(nums, 3)
	fmt.Println("end rotate", nums)
}
