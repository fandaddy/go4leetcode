package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	out := make([]int, 2)
	var j int

	for i, v := range nums {
		for j = i + 1; j < len(nums); j++ {
			if nums[j] == target-v {
				out[0] = i
				out[1] = j
				break
			}
		}
	}

	return out
}

func main() {
	testNums := []int{2, 7, 11, 15}
	out := twoSum(testNums, 9)
	fmt.Println(out)
}
