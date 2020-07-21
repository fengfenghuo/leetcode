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
