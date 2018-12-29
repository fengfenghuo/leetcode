package main

import (
	"fmt"

	"github.com/fengfenghuo/leetcode/RotateArray"
	// "github.com/fengfenghuo/leetcode/RemoveDuplicateNum"
	// "github.com/fengfenghuo/leetcode/InterleavingString"
)

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
