/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans[i-1] = "FizzBuzz"
		} else {
			if i%3 == 0 {
				ans[i-1] = "Fizz"
			} else if i%5 == 0 {
				ans[i-1] = "Buzz"
			} else {
				ans[i-1] = strconv.Itoa(i)
			}
		}
	}
	return ans
}

// @lc code=end

