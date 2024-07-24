<p>You are given an array <code>prices</code> where <code>prices[i]</code> is the price of a given stock on the <code>i<sup>th</sup></code> day.</p>

<p>You want to maximize your profit by choosing a <strong>single day</strong> to buy one stock and choosing a <strong>different day in the future</strong> to sell that stock.</p>

<p>Return <em>the maximum profit you can achieve from this transaction</em>. If you cannot achieve any profit, return <code>0</code>.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> prices = [7,1,5,3,6,4]
<strong>Output:</strong> 5
<strong>Explanation:</strong> Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> prices = [7,6,4,3,1]
<strong>Output:</strong> 0
<strong>Explanation:</strong> In this case, no transactions are done and the max profit = 0.
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li><code>1 &lt;= prices.length &lt;= 10<sup>5</sup></code></li> 
 <li><code>0 &lt;= prices[i] &lt;= 10<sup>4</sup></code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>Array | Dynamic Programming</details><br>

<div>👍 30967, 👎 1161<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价；已支持老用户续费~**



<p><strong><a href="https://labuladong.online/algo/slug.html?slug=best-time-to-buy-and-sell-stock" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

<div id="labuladong_solution_zh">

## 基本思路

**提示：股票系列问题有共通性，但难度较大，初次接触此类问题的话很难看懂下述思路，建议直接看 [详细题解](https://labuladong.online/algo/dynamic-programming/stock-problem-summary/)。**

股票系列问题状态定义：

```python
dp[i][k][0 or 1]
0 <= i <= n - 1, 1 <= k <= K
n 为天数，大 K 为交易数的上限，0 和 1 代表是否持有股票。
```

股票系列问题通用状态转移方程：

```python
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
              max( 今天选择 rest,        今天选择 sell       )

dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
              max( 今天选择 rest,         今天选择 buy         )
```

通用 base case：

```python
dp[-1][...][0] = dp[...][0][0] = 0
dp[-1][...][1] = dp[...][0][1] = -infinity
```

特化到 `k = 1` 的情况，状态转移方程和 base case 如下：

```python
状态转移方程：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], -prices[i])

base case：
dp[i][0] = 0;
dp[i][1] = -prices[i];
```

详细思路解析和空间复杂度优化的解法见详细题解。

**详细题解：[一个方法团灭 LeetCode 股票买卖问题](https://labuladong.online/algo/dynamic-programming/stock-problem-summary/)**

</div>

**标签：[二维动态规划](https://labuladong.online/algo/)，[动态规划](https://labuladong.online/algo/)**

<div id="solution">

## 解法代码



<div class="tab-panel"><div class="tab-nav">
<button data-tab-item="cpp" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">cpp🤖</button>

<button data-tab-item="python" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">python🤖</button>

<button data-tab-item="java" class="tab-nav-button btn active" data-tab-group="default" onclick="switchTab(this)">java🟢</button>

<button data-tab-item="go" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">go🤖</button>

<button data-tab-item="javascript" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">javascript🤖</button>
</div><div class="tab-content">
<div data-tab-item="cpp" class="tab-item " data-tab-group="default"><div class="highlight">

```cpp
// 注意：cpp 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int n = prices.size();
        vector<vector<int>> dp(n, vector<int>(2));
        for (int i = 0; i < n; i++) {
            if (i - 1 == -1) {
                // base case
                dp[i][0] = 0;
                dp[i][1] = -prices[i];
                continue;
            }
            dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
            dp[i][1] = max(dp[i - 1][1], -prices[i]);
        }
        return dp[n - 1][0];
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def maxProfit(self, prices) -> int:
        n = len(prices)
        dp = [[0] * 2 for _ in range(n)]
        for i in range(n):
            if i - 1 == -1:
                # base case
                dp[i][0] = 0
                dp[i][1] = -prices[i]
                continue
            dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
            dp[i][1] = max(dp[i - 1][1], -prices[i])
        return dp[n - 1][0]
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class Solution {
    public int maxProfit(int[] prices) {
        int n = prices.length;
        int[][] dp = new int[n][2];
        for (int i = 0; i < n; i++) {
            if (i - 1 == -1) {
                // base case
                dp[i][0] = 0;
                dp[i][1] = -prices[i];
                continue;
            }
            dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
            dp[i][1] = Math.max(dp[i - 1][1], -prices[i]);
        }
        return dp[n - 1][0];
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func maxProfit(prices []int) int {
    n := len(prices)
    dp := make([][2]int, n)
    for i := 0; i < n; i++ {
        if i-1 == -1 {
            // base case
            dp[i][0] = 0
            dp[i][1] = -prices[i]
            continue
        }
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1], -prices[i])
    }
    return dp[n-1][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var maxProfit = function(prices) {
    let n = prices.length;
    let dp = Array.from({ length: n }, () => Array(2).fill(0));
    for (let i = 0; i < n; i++) {
        if (i - 1 === -1) {
            // base case
            dp[i][0] = 0;
            dp[i][1] = -prices[i];
            continue;
        }
        dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i]);
        dp[i][1] = Math.max(dp[i - 1][1], -prices[i]);
    }
    return dp[n - 1][0];
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🎃🎃 算法可视化 🎃🎃</strong></summary><div id="data_best-time-to-buy-and-sell-stock" data="G/5RIqo3RwCcBXZrx0PqKT7EqRTDvEOarSHIdFKIVuWR1vY4uUc4RewvbatqM1cBb9mNv0r9JcdkNoYf5t5ge3Fk5iRTZ8XwSgb2jF8GqqjdpPTokfU/l6mhzZaFgEyQQ4BSduSCUjpqfenl9O0/E9RRqQAaEMtvpd1La205DmYS+3ABDLwJ0pxF+/2y3uJXKpQmGuvbmfuLLW6l+n3zCC1Cw2SomM4rk7T99xAgQN7KOP+evirxztbv/X10JbG/OvZVvzBwB8PXRhzdkUyqH+JXb8PMjDzDT/mG0Bwdpr5/EQvPt67fyzXX2Z+TuNMaNgeWSW0uJ6a0NxpVyk8TvaGpDHairpqpP3QRMUzf9YsULG9HFWFBYB84nFO/xvB0wz79YLdU7TbDT81y59oPBlrHgRxf8P3R9cgII/tTJuD40+BjfxP046v1agAnZxdHv6rgDJVIjGsvvcs+Voy2cPIA1pwQpnjOal7VIZEPuy8XCEPJGlesCcpGUPD69n9BDkGSzdQ5C3z4l/I33fv70VxhJDdNsmQ6/uOOSXoBcCD9YIZR2TUdR8RTBlu+JNOa5MY3ajcwsvuRa74axRGeBULAYex+YGlbkj8uLZHFZePRcFcmuwj5VaYX/QccGPXxpT6sC0OJwNniZMctA63EchMOlsSc4aOfrbnalHbwMfZclQyl0/3Q5qJwf/93YOlCPHzBP/PQHXkaPPArfrTD67LZwz5sYn4TD2/wzTx8kKfhA7/ihwzTZuN2E0v+EA9v8M08fJCn4QO/4oeM0GaP20N/mB/i4Q2+mYcP8jR84Ff8kAFtdjwPaZm/xMMbPDMPH+xp+OCXncCdbr1ni7gKIzp4opkYvvHQNcfNHybfVFxacb9jad5+6lE5CAte/zVfdmHO9nJdHP4w880qd4wuMvSq5P174ohcLz0jz/kGjQFs8YYIi3/mOS+MQ+BHTKtqYSF7jovant49oFlhXNW1CzwoGCUX0N2tdimsKtI/IHHbeDblUGa1cfFZ4bZj830jxNSy3AcK1e92D1sqWLblRfPg+V3hXaT6GIpbCikLmyTydr28trdAYvQc/7LkuELjA+odytaiu7lMitxgmfaGYSHeRe6jFYff7tuCvGSdajGs5/hPFhwXqeo/yMEtqSVel/o0iYikcgLxjg32ozF4Kt+mjF1c22Zs7CzmEps/DJ6k7Tdi5SXU2iUdsXuhS5w57HvTcULwBccSjlxT5j1ZlIYglcaPxpDGB7hSVqVqJIvLV1BEjVJHciGeHBVVf3pJMqQijRejf1UQkqKiaGRWSclno7P/mTqHH4PK//jiPzyvlkona07czVXlTmxo91AKBpGUrnYtXXYHawF1pVch+GXjAnqBvzGwjT370SveVT6uW9CaLDkeJDbHqPleklI9AJqOCichJlX9rbgJqaplkYt7X4TTslQBIriNAGRECFQ1EwCcwWy9QcoSj1wTOQQ9saFCYfAVbrbqHE2IjddiUp2m4C6qLVUaLy15x17WqwlykVCE3a60DgZByrqIDOGKIZLCxVX1DoMRJu6UVzmI3bXW66JmABRvgUgkJKQ62aIxW5AeXzw74J1QAKscJJJRsAuEuCJg+J+DMfolWnUoHLkmeiJoiKvW12M1l5EFqkIa3U7CIXEExt0pr6OA4nYtXdI92BEYVTuREAZDA66QlEKGQQgUV8olIXqINakMgPoueR4SAqlJtspT3EhPLs6xZ04ogNDaMSMZBauecUVwBrP1hpYsiUeuiRyChrhCDyo5dkLcMG+S2oarwWtyp7bxygjguNqr2WeuJtRYZ0Hv4CeBUbUTGcJgaMAVklq8MAiB4kqlHEQPsSG1AyAwv9GRxjrsIBC+nB4u3oOu3bDt23xeLKVQFTSuCBj+52GMhGCeoiMcuSF6ImiIK/RstTkoJOJWcocZskDcnVo68yNu19Il3YMdgVG1EwlhMDTgCkmnkGEQAsWVoiRED7GQugsPOwEkXqb4ktyOCrshWKUl1t6KWSpid/ASIp1n7czWG+6Ql3BQcEfGUUvT2XSVm9GxNbqdzI+DfIwD4QL8y6iW6saGP0zXKcWTzcg+E7SBVPaks4xchgq12KKIeiMe9hJdU7fpX/Oxj/PTYhwUhyMvF/N8l6XPlUakrRodMPQTizbgV2E4j301vxAv/MH/ZhzHbAd51Gsq7sOoF7iaXjdwHlQ3sGDVD9QARIVjhzvAUTAdcC9OB7ElooOcj7ZVpT9PPf/l/5RjuMqIqTT0mCLZ2sKAg+7bEx7efzjaNCKzGh54b0UOSrkGoHkzthbjbBKLdYVBjHkOi30FUAI5TAoVUEoPw+GlIOFtGCEeJELLjaz3ZI6zs12NYXFWjJrspQPOQFuLLyHknZDknZDBnYgI3gkJ3IlpYfg/YfAoVSARKioThoga4lJhoqFEqGhdMkx0JiEV5htottJbzJZwSclsgUlvoS7hkpDZ4raEY6PZTo31FQ5zjEM2t8vfv4dm2/5rs1NdN4wxNslr1oyEYWaB8OsIx2neImfrz4V7dnrGO7pnzZuLPK02e+L1poWf31/3nn8+XPDztCmrL3ZVT5oHY9kMBv9SI40Lh+yGAcMPdgNjECd3d1aM1X5huderNO8Oe925vKn8GNX+TTyLfmVHLFlrzIY1X+7mSTJy+jKlfau/BlqHXcrTR727vGLNYRvGai0aUYebiVNYkhFUvXnjNaKye/mPolM502xz01Pq/97qHpFUIlvgC52VrD6iQc+b5aYx8KjzlxEmW5Wasf68VhUbrutSAQ=="></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_best-time-to-buy-and-sell-stock"></div></div>
</details><hr /><br />

**类似题目**：
  - [122. 买卖股票的最佳时机 II 🟠](/problems/best-time-to-buy-and-sell-stock-ii)
  - [123. 买卖股票的最佳时机 III 🔴](/problems/best-time-to-buy-and-sell-stock-iii)
  - [188. 买卖股票的最佳时机 IV 🔴](/problems/best-time-to-buy-and-sell-stock-iv)
  - [309. 最佳买卖股票时机含冷冻期 🟠](/problems/best-time-to-buy-and-sell-stock-with-cooldown)
  - [714. 买卖股票的最佳时机含手续费 🟠](/problems/best-time-to-buy-and-sell-stock-with-transaction-fee)
  - [剑指 Offer 63. 股票的最大利润 🟠](/problems/gu-piao-de-zui-da-li-run-lcof)

</div>

</details>
</div>

