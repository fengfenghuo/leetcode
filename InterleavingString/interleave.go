package interleave

// Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

// Example 1:

// Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// Output: true
// Example 2:

// Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// Output: false

func IsInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleave(s1, s2, s3)
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	index1 := 0
	index2 := 0

	for i := 0; i < len(s3); i++ {
		var cha1, cha2 byte
		if index1 >= len(s1) {
			cha1 = 0
		} else {
			cha1 = s1[index1]
		}

		if index2 >= len(s2) {
			cha2 = 0
		} else {
			cha2 = s2[index2]
		}

		if index1 < len(s1) && cha1 == s3[i] {
			index1++
		} else if index2 < len(s2) && cha2 == s3[i] {
			index2++
		} else {
			return false
		}
	}
	return true
}
