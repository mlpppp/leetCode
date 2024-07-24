<p>You are given the heads of two sorted linked lists <code>list1</code> and <code>list2</code>.</p>

<p>Merge the two lists into one <strong>sorted</strong> list. The list should be made by splicing together the nodes of the first two lists.</p>

<p>Return <em>the head of the merged linked list</em>.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg" style="width: 662px; height: 302px;" /> 
<pre>
<strong>Input:</strong> list1 = [1,2,4], list2 = [1,3,4]
<strong>Output:</strong> [1,1,2,3,4,4]
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> list1 = [], list2 = []
<strong>Output:</strong> []
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> list1 = [], list2 = [0]
<strong>Output:</strong> [0]
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li>The number of nodes in both lists is in the range <code>[0, 50]</code>.</li> 
 <li><code>-100 &lt;= Node.val &lt;= 100</code></li> 
 <li>Both <code>list1</code> and <code>list2</code> are sorted in <strong>non-decreasing</strong> order.</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>Linked List | Recursion</details><br>

<div>👍 21786, 👎 2120<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价！算法可视化编辑器上线，[点击体验](https://labuladong.online/algo/intro/visualize/)！**



<p><strong><a href="https://labuladong.online/algo/slug.html?slug=merge-two-sorted-lists" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

> 本文有视频版：[链表双指针技巧全面汇总](https://www.bilibili.com/video/BV1q94y1X7vy)

经典算法题了，[双指针技巧](https://labuladong.online/algo/essential-technique/linked-list-skills-summary/) 用起来。

![](https://labuladong.online/algo/images/链表技巧/1.gif)

这个算法的逻辑类似于「拉拉链」，`l1, l2` 类似于拉链两侧的锯齿，指针 `p` 就好像拉链的拉索，将两个有序链表合并。

**代码中还用到一个链表的算法题中是很常见的「虚拟头结点」技巧，也就是 `dummy` 节点**，它相当于是个占位符，可以避免处理空指针的情况，降低代码的复杂性。

**详细题解：[双指针技巧秒杀七道链表题目](https://labuladong.online/algo/essential-technique/linked-list-skills-summary/)**

**标签：[数据结构](https://labuladong.online/algo/)，[链表](https://labuladong.online/algo/)，[链表双指针](https://labuladong.online/algo/)**

## 解法代码

提示：🟢 标记的是我写的解法代码，🤖 标记的是 chatGPT 翻译的多语言解法代码。如有错误，可以 [点这里](https://github.com/labuladong/fucking-algorithm/issues/1113) 反馈和修正。

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

class LeetCode_Java.List_LinkedList.Solution_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        // 虚拟头结点
        ListNode* dummy = new ListNode(-1);
        ListNode* p = dummy;
        ListNode* p1 = l1;
        ListNode* p2 = l2;

        while (p1 != nullptr && p2 != nullptr) {
            // 比较 p1 和 p2 两个指针
            // 将值较小的的节点接到 p 指针
            if (p1->val > p2->val) {
                p->next = p2;
                p2 = p2->next;
            } else {
                p->next = p1;
                p1 = p1->next;
            }
            // p 指针不断前进
            p = p->next;
        }

        if (p1 != nullptr) {
            p->next = p1;
        }

        if (p2 != nullptr) {
            p->next = p2;
        }

        return dummy->next;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class LeetCode_Java.List_LinkedList.Solution_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        # 虚拟头结点
        dummy = ListNode(-1)
        p = dummy
        p1 = l1
        p2 = l2

        while p1 is not None and p2 is not None: # <extend down -200>![](https://labuladong.online/algo/images/链表技巧/1.gif) #
            # 比较 p1 和 p2 两个指针
            # 将值较小的的节点接到 p 指针
            if p1.val > p2.val:
                p.next = p2
                p2 = p2.next
            else:
                p.next = p1
                p1 = p1.next
            # p 指针不断前进
            p = p.next

        if p1 is not None:
            p.next = p1

        if p2 is not None:
            p.next = p2

        return dummy.next
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class LeetCode_Java.List_LinkedList.Solution_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        // 虚拟头结点
        ListNode dummy = new ListNode(-1), p = dummy;
        ListNode p1 = l1, p2 = l2;

        while (p1 != null && p2 != null) {/**<extend down -200>![](https://labuladong.online/algo/images/链表技巧/1.gif) */
            // 比较 p1 和 p2 两个指针
            // 将值较小的的节点接到 p 指针
            if (p1.val > p2.val) {
                p.next = p2;
                p2 = p2.next;
            } else {
                p.next = p1;
                p1 = p1.next;
            }
            // p 指针不断前进
            p = p.next;
        }

        if (p1 != null) {
            p.next = p1;
        }

        if (p2 != null) {
            p.next = p2;
        }

        return dummy.next;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    // 虚拟头结点
    dummy := &ListNode{Val: -1}
    p := dummy
    p1 := l1
    p2 := l2

    for p1 != nil && p2 != nil {
        // 比较 p1 和 p2 两个指针
        // 将值较小的的节点接到 p 指针
        if p1.Val > p2.Val {
            p.Next = p2
            p2 = p2.Next
        } else {
            p.Next = p1
            p1 = p1.Next
        }
        // p 指针不断前进
        p = p.Next
    }

    if p1 != nil {
        p.Next = p1
    }

    if p2 != nil {
        p.Next = p2
    }

    return dummy.Next
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var mergeTwoLists = function(l1, l2) {
    // 虚拟头结点
    let dummy = new ListNode(-1), p = dummy;
    let p1 = l1, p2 = l2;

    while (p1 !== null && p2 !== null) {
        // 比较 p1 和 p2 两个指针
        // 将值较小的的节点接到 p 指针
        if (p1.val > p2.val) {
            p.next = p2;
            p2 = p2.next;
        } else {
            p.next = p1;
            p1 = p1.next;
        }
        // p 指针不断前进
        p = p.next;
    }

    if (p1 !== null) {
        p.next = p1;
    }

    if (p2 !== null) {
        p.next = p2;
    }

    return dummy.next;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>👾👾 算法可视化 👾👾</strong></summary><div id="data_merge-two-sorted-lists" data="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"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_merge-two-sorted-lists"></div></div>
</details><hr /><br />

**类似题目**：
  - [1305. 两棵二叉搜索树中的所有元素 🟠](/problems/all-elements-in-two-binary-search-trees)
  - [141. 环形链表 🟢](/problems/linked-list-cycle)
  - [142. 环形链表 II 🟠](/problems/linked-list-cycle-ii)
  - [160. 相交链表 🟢](/problems/intersection-of-two-linked-lists)
  - [19. 删除链表的倒数第 N 个结点 🟠](/problems/remove-nth-node-from-end-of-list)
  - [23. 合并K个升序链表 🔴](/problems/merge-k-sorted-lists)
  - [264. 丑数 II 🟠](/problems/ugly-number-ii)
  - [313. 超级丑数 🟠](/problems/super-ugly-number)
  - [86. 分隔链表 🟠](/problems/partition-list)
  - [876. 链表的中间结点 🟢](/problems/middle-of-the-linked-list)
  - [88. 合并两个有序数组 🟢](/problems/merge-sorted-array)
  - [97. 交错字符串 🟠](/problems/interleaving-string)
  - [977. 有序数组的平方 🟢](/problems/squares-of-a-sorted-array)
  - [剑指 Offer 22. 链表中倒数第k个节点 🟢](/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof)
  - [剑指 Offer 25. 合并两个排序的链表 🟢](/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof)
  - [剑指 Offer 49. 丑数 🟠](/problems/chou-shu-lcof)
  - [剑指 Offer 52. 两个链表的第一个公共节点 🟢](/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof)
  - [剑指 Offer II 021. 删除链表的倒数第 n 个结点 🟠](/problems/SLwz0R)
  - [剑指 Offer II 022. 链表中环的入口节点 🟠](/problems/c32eOV)
  - [剑指 Offer II 023. 两个链表的第一个重合节点 🟢](/problems/3u1WK4)
  - [剑指 Offer II 078. 合并排序链表 🔴](/problems/vvXgSW)

</details>
</div>

