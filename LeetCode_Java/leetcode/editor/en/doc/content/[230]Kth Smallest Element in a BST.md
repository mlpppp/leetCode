<p>Given the <code>root</code> of a binary search tree, and an integer <code>k</code>, return <em>the</em> <code>k<sup>th</sup></code> <em>smallest value (<strong>1-indexed</strong>) of all the values of the nodes in the tree</em>.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg" style="width: 212px; height: 301px;" /> 
<pre>
<strong>Input:</strong> root = [3,1,4,null,2], k = 1
<strong>Output:</strong> 1
</pre>

<p><strong class="example">Example 2:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg" style="width: 382px; height: 302px;" /> 
<pre>
<strong>Input:</strong> root = [5,3,6,2,4,null,null,1], k = 3
<strong>Output:</strong> 3
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li>The number of nodes in the tree is <code>n</code>.</li> 
 <li><code>1 &lt;= k &lt;= n &lt;= 10<sup>4</sup></code></li> 
 <li><code>0 &lt;= Node.val &lt;= 10<sup>4</sup></code></li> 
</ul>

<p>&nbsp;</p> 
<p><strong>Follow up:</strong> If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?</p>

<details><summary><strong>Related Topics</strong></summary>Tree | Depth-First Search | Binary Search Tree | Binary Tree</details><br>

<div>👍 11459, 👎 229<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价；已支持老用户续费~**



<p><strong><a href="https://labuladong.online/algo/slug.html?slug=kth-smallest-element-in-a-bst" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

<div id="labuladong_solution_zh">

## 基本思路

BST 的中序遍历结果是有序的（升序），所以用一个外部变量记录中序遍历结果第 `k` 个元素即是第 `k` 小的元素。

**详细题解：[东哥带你刷二叉搜索树（特性篇）](https://labuladong.online/algo/data-structure/bst-part1/)**

</div>

**标签：[二叉搜索树](https://labuladong.online/algo/)，[数据结构](https://labuladong.online/algo/)**

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

class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        // 利用 BST 的中序遍历特性
        traverse(root, k);
        return res;
    }

private:
    // 记录结果
    int res = 0;
    // 记录当前元素的排名
    int rank = 0;
    void traverse(TreeNode* root, int k) {
        if (root == nullptr) {
            return;
        }
        traverse(root->left, k);
        // 中序代码位置
        rank++;
        if (k == rank) {
            // 找到第 k 小的元素
            res = root->val;
            return;
        }

        traverse(root->right, k);
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution:
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        # 利用 BST 的中序遍历特性
        self.traverse(root, k)
        return self.res

    # 记录结果
    res = 0
    # 记录当前元素的排名
    rank = 0

    def traverse(self, root: TreeNode, k: int):
        if root is None:
            return
        self.traverse(root.left, k)
        # 中序代码位置
        self.rank += 1
        if k == self.rank:
            # 找到第 k 小的元素
            self.res = root.val
            return

        self.traverse(root.right, k)
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
    public int kthSmallest(TreeNode root, int k) {
        // 利用 BST 的中序遍历特性
        traverse(root, k);
        return res;
    }

    // 记录结果
    int res = 0;
    // 记录当前元素的排名
    int rank = 0;
    void traverse(TreeNode root, int k) {
        if (root == null) {
            return;
        }
        traverse(root.left, k);
        // 中序代码位置
        rank++;
        if (k == rank) {
            // 找到第 k 小的元素
            res = root.val;
            return;
        }

        traverse(root.right, k);
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func kthSmallest(root *TreeNode, k int) int {
    // 利用 BST 的中序遍历特性
    var res int
    var rank int
    traverse(root, k, &res, &rank)
    return res
}

// 记录结果
// 记录当前元素的排名
func traverse(root *TreeNode, k int, res *int, rank *int) {
    if root == nil {
        return
    }
    traverse(root.Left, k, res, rank)
    // 中序代码位置
    *rank++
    if k == *rank {
        // 找到第 k 小的元素
        *res = root.Val
        return
    }
    traverse(root.Right, k, res, rank)
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var kthSmallest = function(root, k) {
    // 记录结果
    let res = 0;
    // 记录当前元素的排名
    let rank = 0;

    var traverse = function(root, k) {
        if (root === null) {
            return;
        }
        traverse(root.left, k);
        // 中序代码位置
        rank++;
        if (k === rank) {
            // 找到第 k 小的元素
            res = root.val;
            return;
        }
        traverse(root.right, k);
    };

    // 利用 BST 的中序遍历特性
    traverse(root, k);
    return res;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🌟🌟 算法可视化 🌟🌟</strong></summary><div id="data_kth-smallest-element-in-a-bst" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_kth-smallest-element-in-a-bst"></div></div>
</details><hr /><br />

**类似题目**：
  - [1038. 从二叉搜索树到更大和树 🟠](/problems/binary-search-tree-to-greater-sum-tree)
  - [538. 把二叉搜索树转换为累加树 🟠](/problems/convert-bst-to-greater-tree)
  - [剑指 Offer 54. 二叉搜索树的第k大节点 🟢](/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof)
  - [剑指 Offer II 054. 所有大于等于节点的值之和 🟠](/problems/w6cpku)

</div>

</details>
</div>

