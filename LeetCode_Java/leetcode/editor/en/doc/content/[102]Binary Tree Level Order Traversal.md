<p>Given the <code>root</code> of a binary tree, return <em>the level order traversal of its nodes' values</em>. (i.e., from left to right, level by level).</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" style="width: 277px; height: 302px;" /> 
<pre>
<strong>Input:</strong> root = [3,9,20,null,null,15,7]
<strong>Output:</strong> [[3],[9,20],[15,7]]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> root = [1]
<strong>Output:</strong> [[1]]
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> root = []
<strong>Output:</strong> []
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li>The number of nodes in the tree is in the range <code>[0, 2000]</code>.</li> 
 <li><code>-1000 &lt;= Node.val &lt;= 1000</code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>Tree | Breadth-First Search | Binary Tree</details><br>

<div>👍 15317, 👎 315<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价；已支持老用户续费~**

<details><summary><strong>labuladong 思路</strong></summary>

<div id="labuladong_solution_zh">

## 基本思路

前文 [BFS 算法框架](https://labuladong.online/algo/essential-technique/bfs-framework/) 就是由二叉树的层序遍历演变出来的。

下面是层序遍历的一般写法，通过一个 while 循环控制从上向下一层层遍历，for 循环控制每一层从左向右遍历：

![](https://labuladong.online/algo/images/dijkstra/1.jpeg)

</div>

**标签：[BFS 算法](https://labuladong.online/algo/)，[二叉树](https://labuladong.online/algo/)**

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
#include <queue>

class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
public:
    std::vector<std::vector<int>> levelOrder(TreeNode* root) {
        std::vector<std::vector<int>> res;
        if (root == nullptr) {
            return res;
        }

        std::queue<TreeNode*> q;
        q.push(root);
        // while 循环控制从上向下一层层遍历
        while (!q.empty()) {
            int sz = q.size();
            // 记录这一层的节点值
            std::vector<int> level;
            // for 循环控制每一层从左向右遍历
            for (int i = 0; i < sz; i++) {
                TreeNode* cur = q.front();
                q.pop();
                level.push_back(cur->val);
                if (cur->left != nullptr)
                    q.push(cur->left);
                if (cur->right != nullptr)
                    q.push(cur->right);
            }
            res.push_back(level);
        }
        return res;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

from collections import deque

class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        res = []
        if root is None:
            return res

        q = deque()
        q.append(root)
        # while 循环控制从上向下一层层遍历
        while q:
            sz = len(q)
            # 记录这一层的节点值
            level = []
            # for 循环控制每一层从左向右遍历
            for i in range(sz):
                cur = q.popleft()
                level.append(cur.val)
                if cur.left is not None:
                    q.append(cur.left)
                if cur.right is not None:
                    q.append(cur.right)
            res.append(level)
        return res
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> res = new LinkedList<>();
        if (root == null) {
            return res;
        }

        Queue<TreeNode> q = new LinkedList<>();
        q.offer(root);
        // while 循环控制从上向下一层层遍历
        while (!q.isEmpty()) {
            int sz = q.size();
            // 记录这一层的节点值
            List<Integer> level = new LinkedList<>();
            // for 循环控制每一层从左向右遍历
            for (int i = 0; i < sz; i++) {
                TreeNode cur = q.poll();
                level.add(cur.val);
                if (cur.left != null)
                    q.offer(cur.left);
                if (cur.right != null)
                    q.offer(cur.right);
            }
            res.add(level);
        }
        return res;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func levelOrder(root *TreeNode) [][]int {
    var res [][]int
    if root == nil {
        return res
    }

    q := []*TreeNode{root}
    // while 循环控制从上向下一层层遍历
    for len(q) > 0 {
        sz := len(q)
        // 记录这一层的节点值
        level := make([]int, sz)
        // for 循环控制每一层从左向右遍历
        for i := 0; i < sz; i++ {
            cur := q[0]
            q = q[1:]
            level[i] = cur.Val
            if cur.Left != nil {
                q = append(q, cur.Left)
            }
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        }
        res = append(res, level)
    }
    return res
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var levelOrder = function(root) {
    let res = [];
    if (!root) {
        return res;
    }

    let q = [];
    q.push(root);
    // while 循环控制从上向下一层层遍历
    while (q.length > 0) {
        let sz = q.length;
        // 记录这一层的节点值
        let level = [];
        // for 循环控制每一层从左向右遍历
        for (let i = 0; i < sz; i++) {
            let cur = q.shift();
            level.push(cur.val);
            if (cur.left)
                q.push(cur.left);
            if (cur.right)
                q.push(cur.right);
        }
        res.push(level);
    }
    return res;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🌈🌈 算法可视化 🌈🌈</strong></summary><div id="data_binary-tree-level-order-traversal" data="G+1OUZSn0b1EURIoFaCHAttYGvl9lxE8OoqLi+Kxp67bFLyirt8iX3NQYvidkTU1S5dI300bVtWn9tyJLaU7oBr4X7OmMEPJttQNibvfhqSMqWmXS1mqECleml1f3Hc5+Xl6nVvIFQgEGAdKtThnxAkBnSM9/NrmSR6VYLgG4XKjkzspIL7fstk+oXSHkGvcykN5nufyMxtKKxZ3qL7Zoc3RuuPhcAiJMjiBXz9NAx1kOSjDwzxBFwJOl9JhaYVmbqXsTKATFjNw6/dX362ntC4BTuWg/1OlmBBBjf0YU5/nIrNTCTc8QG2K+fbv7P92Zt7Wv2f3XlkT5Wbu2znCwAOoKVAHB8JcX1z/ASJD4Ips9RComz0YY//RF8zN/j0vYs2/TuKZfaIZvpzV8GAOT0961IgoybcstrD+b+f04W+6VRCmviCvDu0e+96JGBNTYEbjwzEcd0zpS1+q5VU4fzYfr2/HUX9BtXUcvf2Cjgeg04PYGnsv9k50/41tu+77M+sJH6fsdG2ViORLWcBKjqdf/+JU+m9vDhXGRtO+9n07yeUyDGb45Jf8RcZCc3SKHvHNc7ptxp3dh9BKQDl6Jq5mrzM8TFKKv4Z6IFYpkNqw7lZehiNPrT76g9xkJan7vpxXKmZmDpRKaUPgMhMsIaTLhPYcKBOnZqbQZQ+DVtPnVwBeQH+si7Ojw715S/zYL8DsdMn2gr45gnUbbnwNhIzKERBDVJ7QtCivp+EKuSFwE6CtIOo8aAbMsEg1SVMf80ofteRjeJiESBzO59/eGbdlQzjrlYoa4hgVSqIX2T2APdw/xWxylU0P+Xi1fZ9JqUUNEj8hRnfUUh/O54ye10x/9UjL9PJlVi3Wqm1mYXXlSUNGq9bTlyewJDjhHXlyNdZVaus2q+ICilpYZihw3ib3pGbejlbpi2wT7yOGPyBlejljo5XG0DoqwGa6w1U6Y01cmux7R7vMWS5w5zaRnzI9KU1gMMUq4bMMIIAeHUnsJM8gy6uh1isZZbEOrNJCR4pvlHcfdcAM2IEuE+l8VLGo1jmkDNdUkyRy535zaWjG90ClW03ZSYXNMFymyDEwaJDqHJK3M4SUAUzBoqA+DIVyqCSXMi7Ndi4djKgLDI3J1ErbdK0uln7RYp4TI21wLgOEiryP1qKamkN7FV72FPz1q1ujO15WDsFU8kKj8mbNnH6gvBVDV3PRNxgp/KgjoTFKlfj5Ta2UlDUWw81KPG1rAv1R7tZR8HLfzKPmJyKFMHAuG6Vs4lxdh0g4gFowZIwGoSVqqDeyrQ7rampjmNkcSFFTSNDsU2SamUnfm/53Ujwlo7Cki5eF+T1g3IylmUtQbdFi840gzzuHDavjDTRahAhErOaNkjoGl70YziijHi2zA5aKlDWqwJIPEdlwhb48iQ9ZJR0imzvWmV5N9jA9sZRPlDjx/7/E3+zEy4U6PRBpi6W0CpnBNsx7sT3KdTO6ZwwabhW2YawjNRThAs15gAtzyUtjbo3NEKFsiSeGEAjaO8WhwnqcC+jVZTJnygsMmLpEi1MZlCUulh4OlrspqUHwHeb3+tYafV5JhFf1+mmJ0tWPH2KhvNHVl44S7rq1fXGQ+JufeJ5HCGVLPTGEoDZKjkNVYs1B6GtI2rZac87rhiKovF5S0AOpJ7/F6OYmwPeblZgEbNKRuS2kU+9kONBve/+5l2OhMfyKc7Fdgraix0m4Pn48P8wrC/r7T78cCY1RqsQXMPWFTX3ISmslrc6l9M7ruhElMba+5Br7sA+JJaC7j8bwWVekxB/zGC4Pp386Yimn8Fn3E9Rt9MgFPD02dmR5rcdNTrwkrUVxE784SPzNTzybI4RvccQQxHqG41BhDmh/AWN1q8G58wIDpi4xFe2rQ3WjTS49XBw3AWUS9jERTd8+fGr3E9Rt9NKXkxIOl/tF0G18eR5B+FaOGIJYz+I4VBjD7jJlzZ00e88/iLohhu6uZlrtS5pYerhYbgJiEPZLrIkY+vbh0Ww6K3biHdI5H9ekTDYSVVKT4e26JuNjzCYyNOeSF0Ehl5aAr9Ef3Mh+3QT4N9PlPEo9Rhuo9S70AIyVfv89ZdPZEYvSqMINWL2AaYHasAPLDg1T0hO1jek6kR2ywUV/3RQ8ZgyaMXy21Tj9HLsq51tVZeS4rhSOWWkrtoxiKZI9Mi2Uzg93t9es0saNvNuOlxhl1cDgIbrmdq53/HzKoFToI6hE9q6c4a5PJjK4fLpbPVzid1cMLqeL06lMFsoVoMYC6ByrWWZJNTwdT9sPQ3+gEAbyzGKMPeUh1kxyMM2pAGJoVS01u+N/d2+V31Pv2H8BxZxxxb65MEnnF+Poj/f4SsE5J6sphQAthqCqDANzKgxVueVQtW0nqaItVMcWVr2KiZfUjv2AQb9I09HV1th/6Tu6/FsVWC0avVl1IDJSkRYTyOUGTCP0+I6YpE6NItCMR5gmEt+SkrvoZCAmcdopQExSp0b/aFw0KfiGzNVOA0ISF50ViElIyV10OpCSu2gJMCaDKUqJ4h0CfpuilKaMogQlpSmnKCGqVzZbrm3Kr8Ddm5kp9O0eKP7CYVhlT2sAbbYyZwsBf6/xXONXJEHZ+73I8yMyfP60uOJpGJ1q/QKy8RXOax5LtxLPPf9fw65FbrXBYXqarP5q0E+95MgP5wB9eMeMa2rLqg6FrDtWu6xvtlnn6C+eo68htBx8dh12Pu4jlJjmxa9x7WMIyAuRppa7z4JlUuQlkdr15HA0gWk5hlCzJgeqaY7eOc+77CZUq8Xk6t+2F6IrHoUERXxgPFQbYuLjGmZFyvGcAjurDQ6/1PdqeaFMIYJiVe7mq3U9VICrHTPL9ZwV8ohYhgWjuTofV69ub8/0PrARaCXZZq3iwJFlhqEfQEMMpxz6K/s2aoPBOjGd2a0mkbM7DXa6OCDmWa4IZ6/xlyBGl+DgPJKpIpBeZj8nVoTUDXs5SKHDoqTjXYHKfyqzfSclCoFCXL44gKr/vaLDVzuHtI6RFdz6teoW3mrBv1he9r/dC2ApGPznu4RpkTU4Pml4CUr/hLikwmQBeopqE9vT4VjuAw=="></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_binary-tree-level-order-traversal"></div></div>
</details><hr /><br />

**类似题目**：
  - [103. 二叉树的锯齿形层序遍历 🟠](/problems/binary-tree-zigzag-level-order-traversal)
  - [107. 二叉树的层序遍历 II 🟠](/problems/binary-tree-level-order-traversal-ii)
  - [1161. 最大层内元素和 🟠](/problems/maximum-level-sum-of-a-binary-tree)
  - [1302. 层数最深叶子节点的和 🟠](/problems/deepest-leaves-sum)
  - [1609. 奇偶树 🟠](/problems/even-odd-tree)
  - [429. N 叉树的层序遍历 🟠](/problems/n-ary-tree-level-order-traversal)
  - [637. 二叉树的层平均值 🟢](/problems/average-of-levels-in-binary-tree)
  - [919. 完全二叉树插入器 🟠](/problems/complete-binary-tree-inserter)
  - [958. 二叉树的完全性检验 🟠](/problems/check-completeness-of-a-binary-tree)
  - [剑指 Offer 32 - I. 从上到下打印二叉树 🟠](/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof)
  - [剑指 Offer 32 - II. 从上到下打印二叉树 II 🟢](/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof)
  - [剑指 Offer 32 - III. 从上到下打印二叉树 III 🟠](/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof)

</div>

</details>
</div>

