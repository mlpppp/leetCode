<p>Given the <code>root</code> of a binary tree, <em>determine if it is a valid binary search tree (BST)</em>.</p>

<p>A <strong>valid BST</strong> is defined as follows:</p>

<ul> 
 <li>The left <span data-keyword="subtree">subtree</span> of a node contains only nodes with keys <strong>less than</strong> the node's key.</li> 
 <li>The right subtree of a node contains only nodes with keys <strong>greater than</strong> the node's key.</li> 
 <li>Both the left and right subtrees must also be binary search trees.</li> 
</ul>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg" style="width: 302px; height: 182px;" /> 
<pre>
<strong>Input:</strong> root = [2,1,3]
<strong>Output:</strong> true
</pre>

<p><strong class="example">Example 2:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg" style="width: 422px; height: 292px;" /> 
<pre>
<strong>Input:</strong> root = [5,1,4,null,null,3,6]
<strong>Output:</strong> false
<strong>Explanation:</strong> The root node's value is 5 but its right child's value is 4.
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li>The number of nodes in the tree is in the range <code>[1, 10<sup>4</sup>]</code>.</li> 
 <li><code>-2<sup>31</sup> &lt;= Node.val &lt;= 2<sup>31</sup> - 1</code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>Tree | Depth-First Search | Binary Search Tree | Binary Tree</details><br>

<div>👍 16815, 👎 1376<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价；已支持老用户续费~**



<p><strong><a href="https://labuladong.online/algo/slug.html?slug=validate-binary-search-tree" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

<div id="labuladong_solution_zh">

## 基本思路

初学者做这题很容易有误区：BST 不是左小右大么，那我只要检查 `root.val > root.left.val` 且 `root.val < root.right.val` 不就行了？

这样是不对的，因为 BST 左小右大的特性是指 `root.val` 要比左子树的所有节点都更大，要比右子树的所有节点都小，你只检查左右两个子节点当然是不够的。

正确解法是通过使用辅助函数，增加函数参数列表，在参数中携带额外信息，将这种约束传递给子树的所有节点，这也是二叉搜索树算法的一个小技巧吧。

**详细题解：[东哥带你刷二叉搜索树（基操篇）](https://labuladong.online/algo/data-structure/bst-part2/)**

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
    bool isValidBST(TreeNode* root) {
        return isValidBST(root, nullptr, nullptr);
    }

    // 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val
    bool isValidBST(TreeNode* root, TreeNode* min, TreeNode* max) {
        // base case
        if (root == nullptr) return true;
        // 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
        if (min != nullptr && root->val <= min->val) return false;
        if (max != nullptr && root->val >= max->val) return false;
        // 限定左子树的最大值是 root.val，右子树的最小值是 root.val
        return isValidBST(root->left, min, root) 
            && isValidBST(root->right, root, max);
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        return self.isValidBST(root, None, None)

    # 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val
    def isValidBST(self, root: TreeNode, minNode: TreeNode = None, maxNode: TreeNode = None) -> bool:
        # base case
        if root is None:
            return True
        # 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
        if minNode is not None and root.val <= minNode.val:
            return False
        if maxNode is not None and root.val >= maxNode.val:
            return False
        # 限定左子树的最大值是 root.val，右子树的最小值是 root.val
        return self.isValidBST(root.left, minNode, root) and self.isValidBST(root.right, root, maxNode)
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
    public boolean isValidBST(TreeNode root) {
        return isValidBST(root, null, null);
    }

    // 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val
    boolean isValidBST(TreeNode root, TreeNode min, TreeNode max) {
        // base case
        if (root == null) return true;
        // 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
        if (min != null && root.val <= min.val) return false;
        if (max != null && root.val >= max.val) return false;
        // 限定左子树的最大值是 root.val，右子树的最小值是 root.val
        return isValidBST(root.left, min, root)
                && isValidBST(root.right, root, max);
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func isValidBST(root *TreeNode) bool {
    return isValidBSTHelper(root, nil, nil)
}

// 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val
func isValidBSTHelper(root, min, max *TreeNode) bool {
    // base case
    if root == nil {
        return true
    }
    // 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
    if min != nil && root.Val <= min.Val {
        return false
    }
    if max != nil && root.Val >= max.Val {
        return false
    }
    // 限定左子树的最大值是 root.val，右子树的最小值是 root.val
    return isValidBSTHelper(root.Left, min, root) && isValidBSTHelper(root.Right, root, max)
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var isValidBST = function(root) {
    // 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val
    function isValidBST(root, min, max) {
        // base case
        if (root === null) return true;
        // 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
        if (min !== null && root.val <= min.val) return false;
        if (max !== null && root.val >= max.val) return false;
        // 限定左子树的最大值是 root.val，右子树的最小值是 root.val
        return isValidBST(root.left, min, root) 
            && isValidBST(root.right, root, max);
    }
    
    return isValidBST(root, null, null);
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🥳🥳 算法可视化 🥳🥳</strong></summary><div id="data_validate-binary-search-tree" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_validate-binary-search-tree"></div></div>
</details><hr /><br />

**类似题目**：
  - [255. 验证前序遍历序列二叉搜索树 🟠](/problems/verify-preorder-sequence-in-binary-search-tree)
  - [450. 删除二叉搜索树中的节点 🟠](/problems/delete-node-in-a-bst)
  - [700. 二叉搜索树中的搜索 🟢](/problems/search-in-a-binary-search-tree)
  - [701. 二叉搜索树中的插入操作 🟠](/problems/insert-into-a-binary-search-tree)

</div>

</details>
</div>

