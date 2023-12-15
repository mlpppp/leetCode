# 标记

- visited

# List

- feature
  - 连续的存储空间
  - 有可能 长度不是动态的，需要处理扩容缩容的问题
  - O(1) 随机访问, Update 简单
  - 扩容和 Insertion/Deletion 麻烦
    - 扩容：O(n), 需要将整个数组复制到新建的另外一个内存区域
    - Insertion/Deletion: O(n), 大量的 element 移动

```go
arr []int
```

### traverse

### 快慢指针

### 滑动窗口

### Binary Search

# LinkList

- feature 与数组全部正好相反
  - 非连续的存储空间，动态长度
  - 不能随机访问
  - 扩容/增加和删除非常容易 O(1)

```go
type ListNode struct {
    val int
    next *ListNode
}
```

traverse

# Hash Table(dictionary):

- operation: insert/delete/search
- average search time O(1)
  - worst case O(n)
- Hash Table vs dictionary
  - dictionary: Generic way to map key-value pair. might have different implementation
  - hash table: it's a specific implementation of dictionary

-----TODO
理解思想即可
基本没有实现/算法需求

# Binary Tree

```go
type TreeNode struct {
    val int
    left *TreeNode
    right *TreeNode
}
```

traverse: 后序遍历

# Tree

```go
type TreeNode struct {
    val int
    children []*TreeNode
}
```

# Graph
