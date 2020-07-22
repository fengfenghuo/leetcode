package common

import "fmt"

/* 3. 无重复字符的最长子串
 */

// LengthOfLongestSubstring ...
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	maxLength := 1
	newStr := []string{}
	for _, str := range s {
		strS := string(str)
		for i, tmp := range newStr {
			// 存在相同重置数组，记录长度
			if tmp == strS {
				fmt.Println(newStr)
				if len(newStr) > maxLength {
					maxLength = len(newStr)
				}
				newStr = newStr[i+1:]
				fmt.Println(newStr, '\n')
				break
			}
		}
		newStr = append(newStr, strS)
	}

	if maxLength > len(newStr) {
		return maxLength
	}
	return len(newStr)
}

/*
4. 编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。

示例：

输入: numbers = [1,2]
输出: [2,1]
提示：

numbers.length == 2
*/

func SwapNumbers(numbers []int) []int {
	// if numbers[0] == 0 {
	// 	numbers[0] = numbers[1]
	// 	numbers[1] = 0
	// } else if numbers[1] == 0 {
	// 	numbers[1] = numbers[0]
	// 	numbers[0] = 0
	// } else {
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	// }

	return numbers
}

/*
1513. 给你一个二进制字符串 s（仅由 '0' 和 '1' 组成的字符串）。

返回所有字符都为 1 的子字符串的数目。

由于答案可能很大，请你将它对 10^9 + 7 取模后返回。



示例 1：

输入：s = "0110111"
输出：9
解释：共有 9 个子字符串仅由 '1' 组成
"1" -> 5 次
"11" -> 3 次
"111" -> 1 次
示例 2：

输入：s = "101"
输出：2
解释：子字符串 "1" 在 s 中共出现 2 次
示例 3：

输入：s = "111111"
输出：21
解释：每个子字符串都仅由 '1' 组成
示例 4：

输入：s = "000"
输出：0


提示：

s[i] == '0' 或 s[i] == '1'
1 <= s.length <= 10^5
*/

func NumSub(s string) int {
	var result uint64
	// uValue := make(map[int]int)
	count := 0
	for _, char := range s {
		if string(char) == "1" {
			count++
		} else if count > 0 {
			// uValue[count]++
			result += uint64((count + 1) * count / 2)
			count = 0
		}
	}

	if count > 0 {
		result += uint64((count + 1) * count / 2)
	}
	// fmt.Println(uValue)
	// for key, value := range uValue {
	// 	result += uint64((key + 1) * key / 2 * value)
	// }

	return int(result % 1000000007)
}

/*
offer11. 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：

输入：[3,4,5,1,2]
输出：1
示例 2：

输入：[2,2,2,0,1]
输出：0
*/

func MinArray(numbers []int) int {
	j := len(numbers) - 1
	if j <= 0 {
		return numbers[0]
	}

	i := 0

	for i < j {
		m := (i + j) / 2
		if numbers[m] > numbers[j] {
			i = m + 1
		} else if numbers[m] < numbers[j] {
			j = m
		} else {
			j--
		}
	}

	return numbers[i]
}

/*
862. 和至少为 K 的最短子数组
返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。

如果没有和至少为 K 的非空子数组，返回 -1 。



示例 1：

输入：A = [1], K = 1
输出：1
示例 2：

输入：A = [1,2], K = 4
输出：-1
示例 3：

输入：A = [2,-1,2], K = 3
输出：3
*/

// func ShortestSubarray(A []int, K int) int {
// 	result := len(A)
// 	start := 0
// 	sum := 0
// 	maxSum := 0
// 	for i := start; i < len(A); i++ {
// 		sum += A[i]
// 		if sum >= K && result > (i-start+1) {
// 			result = (i - start + 1)
// 		}

// 		if (i+1 == len(A) && sum < K) || (sum >= K) {
// 			if sum > maxSum {
// 				maxSum = sum
// 			}
// 			sum = 0
// 			start++
// 			i = start - 1
// 		}
// 	}

// 	if maxSum < K {
// 		return -1
// 	}

// 	return result
// }

func ShortestSubarray(A []int, K int) int {
	i := 0
	j := len(A) - 1
	maxNum := 0
	for i < j {
		m := (i + j) / 2
		if A[m] > A[j] {
			i = m + 1
		} else if A[m] < A[j] {
			j = m
		} else {
			if maxNum < A[j] {
				maxNum = A[j]
			}
			j--
		}
	}
	return i
}
