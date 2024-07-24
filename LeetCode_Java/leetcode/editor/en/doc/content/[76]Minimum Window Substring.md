<p>Given two strings <code>s</code> and <code>t</code> of lengths <code>m</code> and <code>n</code> respectively, return <em>the <strong>minimum window</strong></em> <span data-keyword="substring-nonempty"><strong><em>substring</em></strong></span><em> of </em><code>s</code><em> such that every character in </em><code>t</code><em> (<strong>including duplicates</strong>) is included in the window</em>. If there is no such substring, return <em>the empty string </em><code>""</code>.</p>

<p>The testcases will be generated such that the answer is <strong>unique</strong>.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = "ADOBECODEBANC", t = "ABC"
<strong>Output:</strong> "BANC"
<strong>Explanation:</strong> The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = "a", t = "a"
<strong>Output:</strong> "a"
<strong>Explanation:</strong> The entire string s is the minimum window.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> s = "a", t = "aa"
<strong>Output:</strong> ""
<strong>Explanation:</strong> Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li><code>m == s.length</code></li> 
 <li><code>n == t.length</code></li> 
 <li><code>1 &lt;= m, n &lt;= 10<sup>5</sup></code></li> 
 <li><code>s</code> and <code>t</code> consist of uppercase and lowercase English letters.</li> 
</ul>

<p>&nbsp;</p> 
<p><strong>Follow up:</strong> Could you find an algorithm that runs in <code>O(m + n)</code> time?</p>

<details><summary><strong>Related Topics</strong></summary>Hash Table | String | Sliding Window</details><br>

<div>👍 17889, 👎 735<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价；已支持老用户续费~**



<p><strong><a href="https://labuladong.online/algo/slug.html?slug=minimum-window-substring" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

<div id="labuladong_solution_zh">

## 基本思路

> 本文有视频版：[滑动窗口算法核心模板框架](https://www.bilibili.com/video/BV1AV4y1n7Zt)

这题就是典型的滑动窗口类题目，一般来说难度略高，解法框架如下：

```cpp
// 滑动窗口算法框架
void slidingWindow(string s, string t) {
    unordered_map<char, int> need, window;
    for (char c : t) need[c]++;

    int left = 0, right = 0;
    int valid = 0;
    while (right < s.size()) {
        // c 是将移入窗口的字符
        char c = s[right];
        // 右移窗口
        right++;
        // 进行窗口内数据的一系列更新
        ...

        // debug 输出的位置
        printf("window: [%d, %d)\n", left, right);

        // 判断左侧窗口是否要收缩
        while (window needs shrink) {
            // d 是将移出窗口的字符
            char d = s[left];
            // 左移窗口
            left++;
            // 进行窗口内数据的一系列更新
            ...
        }
    }
}
```

具体的算法原理看详细题解吧，这里写不下。

**详细题解：[我写了首诗，把滑动窗口算法变成了默写题](https://labuladong.online/algo/essential-technique/sliding-window-framework/)**

</div>

**标签：[数组双指针](https://labuladong.online/algo/)，[滑动窗口](https://labuladong.online/algo/)**

<div id="solution">

## 解法代码



<div class="tab-panel"><div class="tab-nav">
<button data-tab-item="cpp" class="tab-nav-button btn active" data-tab-group="default" onclick="switchTab(this)">cpp🟢</button>

<button data-tab-item="python" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">python🤖</button>

<button data-tab-item="java" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">java🤖</button>

<button data-tab-item="go" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">go🤖</button>

<button data-tab-item="javascript" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">javascript🤖</button>
</div><div class="tab-content">
<div data-tab-item="cpp" class="tab-item active" data-tab-group="default"><div class="highlight">

```cpp
class Solution {
    public:
    string minWindow(string s, string t) {
        unordered_map<char, int> need, window;
        for (char c : t) need[c]++;

        int left = 0, right = 0;
        int valid = 0;
        // 记录最小覆盖子串的起始索引及长度
        int start = 0, len = INT_MAX;/**<extend down -200>![](https://labuladong.online/algo/images/slidingwindow/1.png) */
        while (right < s.size()) {
            // c 是将移入窗口的字符
            char c = s[right];
            // 右移窗口
            right++;
            // 进行窗口内数据的一系列更新
            if (need.count(c)) {
                window[c]++;
                if (window[c] == need[c])
                    valid++;
            }

            // 判断左侧窗口是否要收缩
            while (valid == need.size()) {/**<extend down -200>![](https://labuladong.online/algo/images/slidingwindow/2.png) */
                // 在这里更新最小覆盖子串
                if (right - left < len) {
                    start = left;
                    len = right - left;
                }
                // d 是将移出窗口的字符
                char d = s[left];
                // 左移窗口
                left++;
                // 进行窗口内数据的一系列更新
                if (need.count(d)) {
                    if (window[d] == need[d])
                        valid--;
                    window[d]--;
                }
            }/**<extend up -50>![](https://labuladong.online/algo/images/slidingwindow/4.png) */
        }
        // 返回最小覆盖子串
        return len == INT_MAX ?
                "" : s.substr(start, len);
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        from collections import Counter
        need = Counter(t)
        window = Counter()

        left, right = 0, 0
        valid = 0
        # 记录最小覆盖子串的起始索引及长度
        start, length = 0, float('inf') # <extend down -200>![](https://labuladong.online/algo/images/slidingwindow/1.png) #
        while right < len(s):
            # c 是将移入窗口的字符
            c = s[right]
            right += 1
            # 进行窗口内数据的一系列更新
            if c in need:
                window[c] += 1
                if window[c] == need[c]:
                    valid += 1

            # 判断左侧窗口是否要收缩
            while valid == len(need): # <extend down -200>![](https://labuladong.online/algo/images/slidingwindow/2.png) #
                # 在这里更新最小覆盖子串
                if right - left < length:
                    start = left
                    length = right - left
                # d 是将移出窗口的字符
                d = s[left]
                left += 1
                # 进行窗口内数据的一系列更新
                if d in need:
                    if window[d] == need[d]:
                        valid -= 1
                    window[d] -= 1 # <extend up -50>![](https://labuladong.online/algo/images/slidingwindow/4.png) #
        # 返回最小覆盖子串
        return '' if length == float('inf') else s[start:start+length]
```

</div></div>

<div data-tab-item="java" class="tab-item " data-tab-group="default"><div class="highlight">

```java
// 注意：java 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

class Solution {
    public String minWindow(String s, String t) {
        Map<Character, Integer> need = new HashMap<>();
        Map<Character, Integer> window = new HashMap<>();
        for (char c : t.toCharArray()) need.put(c, need.getOrDefault(c, 0) + 1);

        int left = 0, right = 0;
        int valid = 0;
        // 记录最小覆盖子串的起始索引及长度
        int start = 0, len = Integer.MAX_VALUE;/**<extend down -200>![](https://labuladong.online/algo/images/slidingwindow/1.png) */
        while (right < s.length()) {
            // c 是将移入窗口的字符
            char c = s.charAt(right);
            // 右移窗口
            right++;
            // 进行窗口内数据的一系列更新
            if (need.containsKey(c)) {
                window.put(c, window.getOrDefault(c, 0) + 1);
                if (window.get(c).equals(need.get(c)))
                    valid++;
            }

            // 判断左侧窗口是否要收缩
            while (valid == need.size()) {/**<extend down -200>![](https://labuladong.online/algo/images/slidingwindow/2.png) */
                // 在这里更新最小覆盖子串
                if (right - left < len) {
                    start = left;
                    len = right - left;
                }
                // d 是将移出窗口的字符
                char d = s.charAt(left);
                // 左移窗口
                left++;
                // 进行窗口内数据的一系列更新
                if (need.containsKey(d)) {
                    if (window.get(d).equals(need.get(d)))
                        valid--;
                    window.put(d, window.get(d) - 1);
                }
            }/**<extend up -50>![](https://labuladong.online/algo/images/slidingwindow/4.png) */
        }
        // 返回最小覆盖子串
        return len == Integer.MAX_VALUE ? "" : s.substring(start, start + len);
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

import (
    "math"
)

func minWindow(s string, t string) string {
    need := make(map[rune]int)
    window := make(map[rune]int)
    for _, c := range t {
        need[c]++
    }

    left, right := 0, 0
    valid := 0
    // 记录最小覆盖子串的起始索引及长度
    start, minLen := 0, math.MaxInt64
    
    for right < len(s) {
        // c 是将移入窗口的字符
        c := rune(s[right])
        // 右移窗口
        right++
        // 进行窗口内数据的一系列更新
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }

        // 判断左侧窗口是否要收缩
        for valid == len(need) {
            // 在这里更新最小覆盖子串
            if right-left < minLen {
                start = left
                minLen = right - left
            }
            // d 是将移出窗口的字符
            d := rune(s[left])
            // 左移窗口
            left++
            // 进行窗口内数据的一系列更新
            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }
    // 返回最小覆盖子串
    if minLen == math.MaxInt64 {
        return ""
    }
    return s[start:start+minLen]
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 cpp 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var minWindow = function(s, t) {
    //unordered_map<char, int> need, window;
    let need = new Map();
    let window = new Map();
    for (let c of t) need.set(c, (need.get(c) || 0) + 1);

    let left = 0, right = 0;
    let valid = 0;
    // 记录最小覆盖子串的起始索引及长度
    let start = 0, len = Number.MAX_VALUE;

    while (right < s.length) {
        // c 是将移入窗口的字符
        let c = s[right];
        // 右移窗口
        right++;
        // 进行窗口内数据的一系列更新
        if (need.has(c)) {
            window.set(c, (window.get(c) || 0) + 1);
            if (window.get(c) == need.get(c))
                valid++;
        }

        // 判断左侧窗口是否要收缩
        while (valid == need.size) {
            // 在这里更新最小覆盖子串
            if (right - left < len) {
                start = left;
                len = right - left;
            }
            // d 是将移出窗口的字符
            let d = s[left];
            // 左移窗口
            left++;
            // 进行窗口内数据的一系列更新
            if (need.has(d)) {
                if (window.get(d) == need.get(d))
                    valid--;
                window.set(d, window.get(d) - 1);
            }
        }
    }
    // 返回最小覆盖子串
    return len == Number.MAX_VALUE ? "" : s.substr(start, len);
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🥳🥳 算法可视化 🥳🥳</strong></summary><div id="data_minimum-window-substring" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_minimum-window-substring"></div></div>
</details><hr /><br />

**类似题目**：
  - [3. 无重复字符的最长子串 🟠](/problems/longest-substring-without-repeating-characters)
  - [438. 找到字符串中所有字母异位词 🟠](/problems/find-all-anagrams-in-a-string)
  - [567. 字符串的排列 🟠](/problems/permutation-in-string)
  - [剑指 Offer 48. 最长不含重复字符的子字符串 🟠](/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof)
  - [剑指 Offer II 014. 字符串中的变位词 🟠](/problems/MPnaiL)
  - [剑指 Offer II 015. 字符串中的所有变位词 🟠](/problems/VabMRr)
  - [剑指 Offer II 016. 不含重复字符的最长子字符串 🟠](/problems/wtcaE1)
  - [剑指 Offer II 017. 含有所有字符的最短字符串 🔴](/problems/M1oyTv)

</div>

</details>
</div>

