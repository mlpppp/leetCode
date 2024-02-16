/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// (stack). iterate the string, when encounter an opening bracket, push to the stack. When encounter a closing bracket, if the stack top is the corresponding open bracket, remove it, otherwise (empty stack or non-mathcing bracket), return false. At the end the stack should be empty.

func isValid(s string) bool {
	stack := []byte{}
	leftOf := make(map[byte]byte)
	leftOf['}'] = '{'
	leftOf[']'] = '['
	leftOf[')'] = '('

	for i, _ := range s {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		}

		if s[i] == '}' || s[i] == ']' || s[i] == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != leftOf[s[i]] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

