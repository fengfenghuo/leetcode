package main

import (
	"fmt"

	rotate "github.com/fengfenghuo/leetcode/RotateArray"
	// "github.com/fengfenghuo/leetcode/RemoveDuplicateNum"
	// "github.com/fengfenghuo/leetcode/InterleavingString"
	// sort "github.com/fengfenghuo/leetcode/SortTwoArray"
	// "github.com/fengfenghuo/leetcode/common"
)

// func main() {
// 	// 48, 99, 37, 4, -31
// 	// 140

// 	// [56,-21,56,35,-9]
// 	// 61
// 	A := []int{1, 2}
// 	K := 61
// 	num := common.ShortestSubarray(A, K)
// 	fmt.Println(num)
// }

// func main() {
// 	nums := []int{3, 1}
// 	num := common.MinArray(nums)
// 	fmt.Println(num)
// }

// func main() {
// 	str := "111111"
// 	num := common.NumSub(str)
// 	fmt.Println(num)
// }

// func main() {
// 	nums := []int{0, 12345}
// 	newNum := common.SwapNumbers(nums)
// 	fmt.Println(newNum)
// }

// func main() {
// 	// str := ""
// 	// str := " "
// 	// str := "abcabcbb"
// 	str := "dvdf"

// 	maxLength := common.LengthOfLongestSubstring(str)
// 	fmt.Println(maxLength)
// }

// func main() {
// 	// 输入:
// 	// nums1 = [1,2,3,0,0,0], m = 3
// 	// nums2 = [2,5,6],       n = 3

// 	// 输出: [1,2,2,3,5,6]
// 	nums1 := []int{0}
// 	nums2 := []int{1}

// 	sort.Merge(nums1, 0, nums2, 1)
// 	fmt.Println(nums1, nums2)
// }

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate.Rotate(nums, k)
	fmt.Println(nums)
}

// func main() {
// 	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	len := duplicatenum.RemoveDuplicates(nums)
// 	fmt.Println(len, nums)
// }

// func main() {
// 	s1 := "aabcc"
// 	s2 := "dbbca"
// 	s3 := "aadbbcbcac"

// 	s4 := "aadbbbaccc"

// 	fmt.Println("xxxxx", interleave.IsInterleave(s1, s2, s3))
// 	fmt.Println("xxxxx", interleave.IsInterleave(s1, s2, s4))
// }
