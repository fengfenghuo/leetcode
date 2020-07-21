package calc

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for index, num := range nums {
		for i := index + 1; i < len(nums); i++ {
			if num+nums[i] == target {
				result[0] = index
				result[1] = i
				return result
			}
		}
	}
	return result
}
