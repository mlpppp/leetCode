
//leetcode submit region begin(Prohibit modification and deletion)
class Solution {
    // sliding window
    public int maxProfit(int[] prices) {
        int left = 0, right = 0;
        int maxProfit = 0;
        while(left <= right && right < prices.length) {
            int curProfit = prices[right] - prices[left];
            maxProfit = Math.max(maxProfit, curProfit);
            right++;
            while(left <= right && right < prices.length) {
                curProfit = prices[right] - prices[left];
                if (curProfit < 0) {
                    left++;
                } else {
                    break;
                }
            }
        }
        return maxProfit;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
