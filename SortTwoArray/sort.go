package sort

import (
	"fmt"
)

// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

// 说明:

// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

// 示例:

// 输入:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// 输出: [1,2,2,3,5,6]

func Merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := 0
	for i := 0; i < n; {
		num2 := nums2[i]
		fmt.Println(num2)
		hasSort := false
		for index1 < m+n {
			num1 := nums1[index1]
			if num2 < num1 {
				nums1[index1] = num2
				nums2[i] = num1
				hasSort = true
				index1++
				break
			} else if num1 == 0 && index1 >= m {
				nums1[index1] = num2
				break
			}
		}
		if !hasSort {
			i++
		}
		fmt.Println(i)
	}
}
