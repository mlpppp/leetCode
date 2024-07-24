<p>Given the <code>head</code> of a singly linked list, reverse the list, and return <em>the reversed list</em>.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg" style="width: 542px; height: 222px;" /> 
<pre>
<strong>Input:</strong> head = [1,2,3,4,5]
<strong>Output:</strong> [5,4,3,2,1]
</pre>

<p><strong class="example">Example 2:</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg" style="width: 182px; height: 222px;" /> 
<pre>
<strong>Input:</strong> head = [1,2]
<strong>Output:</strong> [2,1]
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> head = []
<strong>Output:</strong> []
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li>The number of nodes in the list is the range <code>[0, 5000]</code>.</li> 
 <li><code>-5000 &lt;= Node.val &lt;= 5000</code></li> 
</ul>

<p>&nbsp;</p> 
<p><strong>Follow up:</strong> A linked list can be reversed either iteratively or recursively. Could you implement both?</p>

<details><summary><strong>Related Topics</strong></summary>Linked List | Recursion</details><br>

<div>👍 21543, 👎 437<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价！算法可视化编辑器上线，[点击体验](https://labuladong.online/algo/intro/visualize/)！**



<p><strong><a href="https://labuladong.online/algo/slug.html?slug=reverse-linked-list" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

递归实现反转链表常常用来考察递归思想，我这里就用纯递归来翻转链表。

**对于递归算法，最重要的就是明确递归函数的定义**。具体来说，我们的 `reverse` 函数定义是这样的：

**输入一个节点 `head`，将「以 `head` 为起点」的链表反转，并返回反转之后的头结点**。

明白了函数的定义，再来看这个问题。比如说我们想反转这个链表：

![](https://labuladong.online/algo/images/反转链表/1.jpg)

那么输入 `reverse(head)` 后，会在这里进行递归：

```java
ListNode last = reverse(head.next);
```

不要跳进递归（你的脑袋能压几个栈呀？），而是要根据刚才的函数定义，来弄清楚这段代码会产生什么结果：

![](https://labuladong.online/algo/images/反转链表/2.jpg)

这个 `reverse(head.next)` 执行完成后，整个链表就成了这样：

![](https://labuladong.online/algo/images/反转链表/3.jpg)

并且根据函数定义，`reverse` 函数会返回反转之后的头结点，我们用变量 `last` 接收了。

现在再来看下面的代码：

```java
head.next.next = head;
```

![](https://labuladong.online/algo/images/反转链表/4.jpg)

接下来：

```java
head.next = null;
return last;
```

![](https://labuladong.online/algo/images/反转链表/5.jpg)

神不神奇，这样整个链表就反转过来了！

**详细题解：[递归魔法：反转单链表](https://labuladong.online/algo/data-structure/reverse-linked-list-recursion/)**

**标签：单链表**

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
    ListNode* reverseList(ListNode* head) {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        ListNode* last = reverseList(head->next);
        head->next->next = head;
        head->next = nullptr;
        return last;
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class LeetCode_Java.List_LinkedList.Solution_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution:
    # public ListNode reverseList(ListNode head) {
    def reverseList(self, head: ListNode) -> ListNode:
        if head is None or head.next is None:
            return head
        last = self.reverseList(head.next) # <extend up -200>![](https://labuladong.online/algo/images/反转链表/3.jpg) #
        head.next.next = head # <extend up -200>![](https://labuladong.online/algo/images/反转链表/4.jpg) #
        head.next = None # <extend up -200>![](https://labuladong.online/algo/images/反转链表/5.jpg) #
        return last
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class LeetCode_Java.List_LinkedList.Solution_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.LeetCode_Java.List_LinkedList.LeetCode_Java.List_LinkedList.LeetCode_Java.BinaryTree.Solution {
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode last = reverseList(head.next);/**<extend up -200>![](https://labuladong.online/algo/images/反转链表/3.jpg) */
        head.next.next = head;/**<extend up -200>![](https://labuladong.online/algo/images/反转链表/4.jpg) */
        head.next = null;/**<extend up -200>![](https://labuladong.online/algo/images/反转链表/5.jpg) */
        return last;
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    last := reverseList(head.Next)
    // ![](https://labuladong.online/algo/images/反转链表/3.jpg)
    head.Next.Next = head
    // ![](https://labuladong.online/algo/images/反转链表/4.jpg)
    head.Next = nil
    // ![](https://labuladong.online/algo/images/反转链表/5.jpg)
    return last
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var reverseList = function(head) {
    if (head === null || head.next === null) {
        return head;
    }
    var last = reverseList(head.next);
    head.next.next = head;
    head.next = null;
    return last;
};
```

</div></div>
</div></div>

<hr /><details open hint-container details><summary style="font-size: medium"><strong>🌈🌈 算法可视化 🌈🌈</strong></summary><div id="data_reverse-linked-list" data="G5cmACwLeMNo/C1IKggeJW/qfkzNdY6pHECQxobKchWK5FOAKdsaPUMXVaX1OZvT6SUngN2U+XZAeGVn1d5s9QC6oAY3D8e4BNnpySGiqXKi8ou118jw/YSFzUSiMLFUMhZw/Nu9vhKBq7q9DaKqi4wroPEVUpR3om6qVGyMpveVK/8BkbVb36pvIfJuYHnTQxNPsLyFbxi4rFwQdWAhifUTZztYiMnCr6T329Xy4vRmduiDslHm43weFahm3yexrHHxmV+jWvzT6EZqpyFxlQCWpSZ2eeG4/jeNVcCmgr01fVYv51N2MGTms/CZUNjMmy1GQWY7m0porLemm7Dlvdyj5mSqzqZ4U8c5fErJGydsv5nn8W3VIr+XXYSt0JPm4gnK7dtZGJqiscDpYWxubeZd3tBy0J9cMKUlYMX6XLqu9zLLBRVB3hWcsQ1bNkXJfaqhH2adSkt8QpbDsKeeTgaX0ewyPFbXF1Jl45f6JKbHf0dCPR0QPxxdF8YKzArbg2vyvKUKwNT1Z9IUzsLm4QPNZVJgtnoNjgSX6B5bCSul/qSOO9dzX1T54wEq8xC7r5fbU1PK4WMrHDprVcXuVceOC9ha/Amj8L1qX2kuYtv19aKUw8fWOH0OOse2i327gI8Wf8YIb+crzUVsXF9fSjl8bI2zPd+VsXHFFpYFFOnxZ9yHx/lKcxE7btnikmL6cPU7/CFInGqb3WduTeaNIY1l4RuwfGKDnPKNAptdrrr12ywfdeP3rSrPyj+CEVLTMbybMz4uLkCa4X5ixmLWuLMjEW4g6y179iVsi/F1XsFMG5++vUcRBNHEncv3szg6elFYa42+TbuK9cmJldjgbED80UtNpwulUrIjczox8kH5ORyzqLfWUG4gum7rMh/bcWT7f/SqC9nr4vjwx0CXnVLY0c+E3IjAUWYfRr9+Y8F71WuIXUjffjR8BLotlcKOfibkRgZ9txL+RURwuzUEF5K7O4aPSI+lUtjRz0BuZNB3lPAvImaee6PrAqdL8ydWpyRRpUkRCCPXEG8I2bNxgng8tSez2jklV+uAGYz97rHHISdAwQPBgto/MCxoZQUVekUqjIrWFBVo5v6BLehQnMMtmu5L7c0ah7kkaVZOksQGTkaeJxoktACGF8Bjt0XiFoPJAR6bLBIlyjkDPUKNS1amHrAgIhrRxTQa8Di6em6bY947a5AFGWPo5Yl5xIvqAZ5M+0ibLwPveFKFFhTchwTf4hi1KXkeHg0iWI4bgWF5R3QPkrIMrVhWeBtEEKu0sKfWPadymVhttmiUaG206FhrLQbWWoOJCDOSCokmviBzHOKRcO80ThKq6bujOPT5fOp4vVJZlhGKtzCHL5X34T/fC4fVXGZ9l8NG5v0MXSShYX8IJsJ8ISbox2cidh1QQo/deT0M7TTxeGl16L/znox4DMihYfrboz/5DtLHTKf+mlosW5IGUodlT7Ltg6gBDRuqSy3IobdJ+FKd+v/ZiPT04UqykfLGokd8ryPvBvOT+zFXI8MxhupRdiUemjIzFI0QUqkfmYMJ0EZbPF7e1/dHjeVU+5Sjfzif7xmBAQh0G7D+grCQDE3MRw48dH0yXjk5PwEPaeI52BrOGO8nN5hdrwtYsjUK"></div><div class="resizable aspect-ratio-container" style="height: 100%;">
<div id="iframe_reverse-linked-list"></div></div>
</details><hr /><br />

**类似题目**：
  - [92. 反转链表 II 🟠](/problems/reverse-linked-list-ii)
  - [剑指 Offer 24. 反转链表 🟢](/problems/fan-zhuan-lian-biao-lcof)
  - [剑指 Offer II 024. 反转链表 🟢](/problems/UHnkqh)

</details>
</div>

