/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start

// rule of vaild parentheses: 1. #( == #) 2. for a vaild parentheses string str, #( > #) must be true for any i in the substring str[:i]

// follow the rule, we DFS from backtracking(lRemain=n, rRemain=n) and exit when backtracking(lRemain=0, rRemain=0). if remaining #( == #), we must add left '(', otherwise we can either add left or right

func generateParenthesis(n int) []string {
	ans := []string{}
	backtracking(n, n, &[]byte{}, &ans)
	return ans
}

func backtracking(lRemain, rRemain int, str *[]byte, ans *[]string) {
	if lRemain == 0 && rRemain == 0 {
		*ans = append(*ans, string(*str))
		return
	}
	// when #( == #), we must add '('
	if lRemain == rRemain {
		// choose '('
		*str = append(*str, '(')
		backtracking(lRemain-1, rRemain, str, ans)
		*str = (*str)[:len(*str)-1]
	} else {
		// choose '('
		if lRemain != 0 {
			*str = append(*str, '(')
			backtracking(lRemain-1, rRemain, str, ans)
			*str = (*str)[:len(*str)-1]

		}

		// choose ')'
		if rRemain != 0 {
			*str = append(*str, ')')
			backtracking(lRemain, rRemain-1, str, ans)
			*str = (*str)[:len(*str)-1]
		}
	}
}

// @lc code=end

