<p>Given an array of integers <code>nums</code>&nbsp;and an integer <code>target</code>, return <em>indices of the two numbers such that they add up to <code>target</code></em>.</p>

<p>You may assume that each input would have <strong><em>exactly</em> one solution</strong>, and you may not use the <em>same</em> element twice.</p>

<p>You can return the answer in any order.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,7,11,15], target = 9
<strong>Output:</strong> [0,1]
<strong>Explanation:</strong> Because nums[0] + nums[1] == 9, we return [0, 1].
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,2,4], target = 6
<strong>Output:</strong> [1,2]
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,3], target = 6
<strong>Output:</strong> [0,1]
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li><code>2 &lt;= nums.length &lt;= 10<sup>4</sup></code></li> 
 <li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li> 
 <li><code>-10<sup>9</sup> &lt;= target &lt;= 10<sup>9</sup></code></li> 
 <li><strong>Only one valid answer exists.</strong></li> 
</ul>

<p>&nbsp;</p> 
<strong>Follow-up:&nbsp;</strong>Can you come up with an algorithm that is less than 
<code>O(n<sup>2</sup>)</code>
<font face="monospace">&nbsp;</font>time complexity?

<details><summary><strong>Related Topics</strong></summary>Array | Hash Table</details><br>

<div>👍 57071, 👎 1986<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价；已支持老用户续费~**



<p><strong><a href="https://labuladong.online/algo/slug.html?slug=two-sum" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

<div id="labuladong_solution_zh">

## 基本思路

大家都喜欢幽默的人，如果你想调侃自己经常拖延，可以这样调侃下自己（手动狗头）：

背单词背了半年还是 abandon, abandon，刷题刷了半年还是 two sum, two sum...

言归正传，这道题不难，但由于它是 LeetCode 第一题，所以名气比较大，解决这道题也可以有多种思路，我这里说两种最常见的思路。

第一种思路就是靠排序，把 `nums` 排序之后就可以用 [数组双指针技巧汇总](https://labuladong.online/algo/essential-technique/array-two-pointers-summary/) 中讲到的左右指针来求出和为 `target` 的两个数。

不过因为题目要求我们返回元素的索引，而排序会破坏元素的原始索引，所以要记录值和原始索引的映射。

进一步，如果题目拓展延伸一下，让你求三数之和、四数之和，你依然可以用双指针技巧，我在 [一个函数秒杀 nSum 问题](https://labuladong.online/algo/practice-in-action/nsum/) 中写一个函数来解决所有 N 数之和问题。

第二种思路是用哈希表辅助判断。对于一个元素 `nums[i]`，你想知道有没有另一个元素 `nums[j]` 的值为 `target - nums[i]`，这很简单，我们用一个哈希表记录每个元素的值到索引的映射，这样就能快速判断数组中是否有一个值为 `target - nums[i]` 的元素了。

简单说，数组其实可以理解为一个「索引 -> 值」的哈希表映射，而我们又建立一个「值 -> 索引」的映射即可完成此题。

**详细题解：[一个方法团灭 nSum 问题](https://labuladong.online/algo/practice-in-action/nsum/)**

</div>

**标签：[双指针](https://labuladong.online/algo/)，哈希表，[数组](https://labuladong.online/algo/)**

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

class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        // 维护 val -> index 的映射
        unordered_map<int, int> valToIndex;
        for (int i = 0; i < nums.size(); i++) {
            // 查表，看看是否有能和 nums[i] 凑出 target 的元素
            int need = target - nums[i];
            if (valToIndex.find(need) != valToIndex.end()) {
                return {valToIndex[need], i};
            }
            // 存入 val -> index 的映射
            valToIndex[nums[i]] = i;
        }
        return {};
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # 维护 val -> index 的映射
        val_to_index = {}
        for i in range(len(nums)):
            # 查表，看看是否有能和 nums[i] 凑出 target 的元素
            need = target - nums[i]
            if need in val_to_index:
                return [val_to_index[need], i]
            # 存入 val -> index 的映射
            val_to_index[nums[i]] = i
        return []
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
    public int[] twoSum(int[] nums, int target) {
        // 维护 val -> index 的映射
        HashMap<Integer, Integer> valToIndex = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            // 查表，看看是否有能和 nums[i] 凑出 target 的元素
            int need = target - nums[i];
            if (valToIndex.containsKey(need)) {
                return new int[]{valToIndex.get(need), i};
            }
            // 存入 val -> index 的映射
            valToIndex.put(nums[i], i);
        }
        return null;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func twoSum(nums []int, target int) []int {
    // 维护 val -> index 的映射
    valToIndex := make(map[int]int)
    for i, num := range nums {
        // 查表，看看是否有能和 nums[i] 凑出 target 的元素
        need := target - num
        if j, ok := valToIndex[need]; ok {
            return []int{j, i}
        }
        // 存入 val -> index 的映射
        valToIndex[num] = i
    }
    return nil
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var twoSum = function(nums, target) {
    // 维护 val -> index 的映射
    let valToIndex = new Map();
    for (let i = 0; i < nums.length; i++) {
        // 查表，看看是否有能和 nums[i] 凑出 target 的元素
        let need = target - nums[i];
        if (valToIndex.has(need)) {
            return [valToIndex.get(need), i];
        }
        // 存入 val -> index 的映射
        valToIndex.set(nums[i], i);
    }
    return null;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>👾👾 算法可视化 👾👾</strong></summary><div id="data_two-sum" data="G5EeUZSHzS0AdCbYMdFAtuF0N26X0xJRj9EXXaKtWqUX9AmyPTdq/l4NGMyknKi63Pz8/ds2/lARwaNLxKEjaDZakd+7Xy8QCvPPV9RW2LwSSLTEtjuT7C1my+CIbdIpICgkYYCsqXYrV+mq7D6GE94lJa7Q17b+cewzb8vYvxj2ZInYbMXY/8PAc7DIsTnoRZzrS/z0MBPjFx+wPro7O9xbLN+vj/5DmX1LOvaLo+p+T+JlY09m/kVW8/ojXd2e7lwfM2Tdsox3rR/7k6RPHNiBQO9cXx3WcrbsJqrBqswBQpq/YtgswOl9lcGphO0B69N8czkvn5vTVdkz4hO+/PVxs9Y72PhfMUFSnp9b/v0Th/UHV7xBRtL+gLE/uT6+nOtcvYaxz49Ll2k9S+U5a2bmJIqPzCAvcxcviehQbYTjZ5SwZERpoOoJO45ljvjSlUnRDVGCmkAVUtF3bfbdYDqAzAE/23Z1v6g5C/Fc51iZ6LI/p/X6jltVSFWFMrwksmPyt4Q+yvJ+waEjzWJSUqzpzPusjlmlKsUhlE7zlu0A9BIFsnIZ9ukXxUZjSHBizDTS6GicWbGFdV8/J6j7CkL50H0HozP89lTQ4V0UyJF5ds/3VtFsMtLoaCMp0wTPUwOD19dm5VZSqTclWE04QIoqCcOn8J5eozLPRwId9UmRHQOjr+FQD5gFjVTzaXrEwInZFdfAMLojlF3gpStYnuaxSrUqoufFx2ssuWvhu/E0YpzZeOn7ne8d/pTfn/0p1nLBx7nKdAiLKvwHPVJluDlrWEdgAHDyYzYv+aqmKcXMUUHVgLKE3AySKKfkI/DkoYqpDxBYi5ASgYoZYDAJwjBoMIIDQEGx3EALqJKmkjTvCVQJqA9ZKEAyoYzAk4YqpT5CYKNnnBVUygBpnqEkGDQUoQHhyW9J6pyncppqVTImUDmg0DzIDMmEMgJPGqqc+gSBhSyzT6ByBqgpcMswaDjCAyBBQcIIBCqlKW0hAoJKs0Wjucx80rJ7ZMnxe5dSUrXVSekbg3YvIZbQsifXcajexGRDiHmUZyh6jFhlfwH9otCoGFet4+6CrpJsXv2nxp/c6cwyXz87XqlSy2Uqumy5l1dip9rBI1+bkQGUCR3QOI64AfwTQlWhjChBp5BrM0lMLke+depj14UfsX5yjYWxO4suaKnfYbmLo4MzsfVtsjjvffTF4+JkAfFFuDiN02/AZTE8C4UypCff1x4dIRpS9MIOrccTLz67ml4IbXjPTd/5KchGuOF3OltBQLw+sZjHTMQmiebjyaNPRrTyhKHOWGJ0pyKuAlpWakVoRKjkSqmIjYSVVBEa2VdCSfAUnb856UdxMtaZooK6KPntZihWssgF4f+62T2O40nXdd3z593D3ou2DAfcpKNu9abeztf+Jxbt4Mur2km8aYJ5t3gvvUn1yFG76qIGJvY49/d+/cbmy1fxn0dAL3VqPZCwCD3rKK1J4Qqtf2j3ztn/WmNAB/bXV/JLvJe50vnGsD91/9jw05y06gXBbruacTMG/uyBIy5b6/bn9nEr0kmFZvdbvGy4vD6u5x3n6/3++3xwY4tQFSn7Wia7Snfqyz1Fix3DbX+2ly43UtpNU8TXplHCqUOKtt12dQXXUT7LYkb9C8yvnVa+e1bi6O9x7JHMlar40mgaO1ZsapmgB3UN6On25eU4npC9dGeTjuzPe28MK3ArvCqiIj5QJtetbLPEmygyL/jBzhfXJ1f7x20JxJWJfWf/chgIo1Ob3BDIAeid/dGDQ/bLThwdMXdmQyfemgRClQTp8bXT0gM="></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_two-sum"></div></div>
</details><hr /><br />

**类似题目**：
  - [15. 三数之和 🟠](/problems/3sum)
  - [167. 两数之和 II - 输入有序数组 🟠](/problems/two-sum-ii-input-array-is-sorted)
  - [18. 四数之和 🟠](/problems/4sum)
  - [剑指 Offer 57. 和为s的两个数字 🟢](/problems/he-wei-sde-liang-ge-shu-zi-lcof)
  - [剑指 Offer II 007. 数组中和为 0 的三个数 🟠](/problems/1fGaJU)

</div>

</details>
</div>



J/fr+UOl9zxz+vVrwwJhQRtFXCR+naBb+PViobOtspF7Hj+9OEysZu/zrtM54ux/DoMQWJlgQAmAftd6mAhYjk9zDw=="></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_two-sum"></div></div>
</details><hr /><br />

**类似题目**：
  - [15. 三数之和 🟠](/problems/3sum)
  - [167. 两数之和 II - 输入有序数组 🟠](/problems/two-sum-ii-input-array-is-sorted)
  - [18. 四数之和 🟠](/problems/4sum)
  - [剑指 Offer 57. 和为s的两个数字 🟢](/problems/he-wei-sde-liang-ge-shu-zi-lcof)
  - [剑指 Offer II 007. 数组中和为 0 的三个数 🟠](/problems/1fGaJU)

</details>
</div>

