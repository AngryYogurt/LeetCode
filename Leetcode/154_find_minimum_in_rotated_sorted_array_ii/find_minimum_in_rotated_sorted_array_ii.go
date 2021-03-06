package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(1, findMin([]int{1, 1}))
	tools.AssertObjectEqual(1, findMin([]int{10, 1, 10, 10, 10, 10}))
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
 */

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	rotIdx := findRotateIndex(nums)
	return nums[rotIdx]
}

func findRotateIndex(nums []int) int {
	start, end := 0, len(nums)-1
	if nums[end] > nums[start] {
		return 0
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		} else if nums[mid] < nums[start] {
			end = mid
		} else if nums[mid+1] > nums[end] {
			start = mid + 1
		} else {
			end--
		}
	}
	return start
}
