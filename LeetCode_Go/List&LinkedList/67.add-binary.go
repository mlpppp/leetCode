/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start

// Given two binary strings a and b, return their sum as a binary string.

// https://www.youtube.com/watch?v=keuWJ47xG8g

// iterate from the end of two string till the head of the longer string, calculate bit position value and the carry value like basic arithematics. note: be careful when converting int, byte and string.

func addBinary(a string, b string) string {
	ans := make([]byte, 1+max(len(a), len(b)))
	carry := 0 // 0 or 1
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; {
		var aDigit, bDigit int // 0 or 1
		if i >= 0 {
			aDigit = int(a[i] - '0')
		} else {
			aDigit = 0
		}
		if j >= 0 {
			bDigit = int(b[j] - '0')
		} else {
			bDigit = 0
		}
		total := aDigit + bDigit + carry // int: 0, 1, 2, or 3

		// calculate answer's bit position
		char := byte('0' + total%2)
		ans[max(i, j)+1] = char

		// calculate next carry bit
		carry = total / 2 // 0, or 1

		i--
		j--
	}
	// add last carry bit
	if carry == 1 {
		ans[0] = '1'
	} else {
		ans = ans[1:]
	}
	return string(ans)
}

// @lc code=end

