
# Heap
- heap: tree with constrain of: 
  - max heap:
    - Constraints: parent > child 
    - 堆顶的元素必须是数组中的argmax
  - min heap:
    - Constraints: child > parent  
    - 堆顶的元素必须是数组中的argmin
- binary heap: heap but complete binary tree

- implementation dataStruct: list  
![](https://imgur.com/MEFttfJ.jpg)
    - 这样储存的好处是任何一个节点都可以通过数组的index计算得到(因为是完全二叉树)
    - 例如: 
      - 1 的children: 1\*2， 1\*2+1；
      - 3 的children: 3\*2， 3\*2+1
## functions:
- maxHeapify(): O(logn) 通过下沉根节点维护maxHeap的性质
- changeHeapValue(k, val):O(logn) 改变任意节点k的值，同时维护maxHeap的性质
- buildMaxHeap():O(n) 从无序数组中建立maxHeap

### maxHeapify(), O(logn):
- input: 根节点为i的binary tree
- 假定根节点为i, 根节点不满足Heap constraint 
- 假定根节点的left Subtree 和 right Subtree都是maxHeap,
- 目的: 让根节点i下沉到合适的位置使整个二叉树满足Heap constraint 
- 算法
```py
def maxHeapify(Btree(i)):

  l = child(i)[0]
  r = child(i)[1]  # Btree(l)和Btree(r)都应该是heap
  if (i > l and i > r): return 
  else: 
    swap (i, argmax(l, r))  # 将根结点与左右子树中最大的一个树根交换(如果i自己并不是最大的)
    maxHeapify(Btree(i))  # 在交换后的以i为根结点的子树上递归调用maxHeapify (这个子树heap性质可能被破坏)

```

### buildMaxHeap():O(n)
- input: random array of size n. A:n
- 算法：
```py
def buildMaxHeap(A):
  for i = floor(n/2) down to 1:
    maxHeapify(Btree(i))
```
### changeHeapValue(k, val):O(logn)
- input: 根节点为i的heap A, 节点k以及要更改的值val
- output: 更改K的值并且维护heap的性质
- 假定val总是非减的，即总是正数
- 因为heap中结点K增大了需要将结点K上浮到正确的位置
```py
def changeHeapValue(A, k, val):
  A[k] = A[k]+val
  parent = parent(k)
  while A[parent] < A[k]:
    swap(A[k], A[parent])
    parent = parent(k)
```
## application:
- heapSort(): O(nlogn)时间原址排序算法
- priorityQueue: 数据结构
### heapSort(): O(nlogn)
- input: random unsorted array of size n. A:n
- output: 用最大堆实现原地址**升序**排序
- 算法：
 ```py
def heapSort(A):
  buildMaxHeap(A)
  for i = n down to 2:
    swap(A[1], A[i])
    excludeFromHeap(A[i]) # 把堆顶最大的元素移到Array底并且不在之后的堆算法中使用
    maxHeapify(Btree(1)) # 重新使堆顶元素成为剩余元素中的argMax
```
# priorityQueue
队列储存一定的成员(array A of size n), 出队列时出队列的成员(pq[1])一定是数值最大的一个(argmax)
- implement: max binary heap
## methods:
- insertOne(): 需要保证insert以后 pq(堆)顶的元素必须是数组中的argmax
- pollOne(delMax): 需要保证poll以后pq(堆)顶的元素依然是数组中的argmax

### insertOne():O(logn)
- input: heap A of size n,  val: new node's value
- 算法：
```py
def insertOne(A, k, val):
  A.size = n+1
  A[n+1] = -1
  changeHeapValue(A, n+1, val)
```
### pollOne(aka. delMax): O(logn) => maxHeapify()
- input: heap A of size n
- output: i as argmax(A), heap A of size n-1
把堆顶元素i和堆底元素j对调，删除i(出队列)并且对BTree(j)调用maxHeapify()
- 算法：
```py
def pollOne(A):
  i = A[1]
  j = A[n]
  swap(A[1], A[n])
  del A[n]
  maxHeapify(Btree(1))

  return i
```

# BST: binary search tree
### what is it
- a binary tree
  - key, left, right, parent
    - key: the value
    - left, right, parent: all pointers
- satisfy the **binary-search-tree property**:
  - Let $x$ be a node in a binary search tree. If $y$ is a node in the left subtree of $x$, then $y.key\leq x.key$. If $y$ is a node in the right subtree of $x$, then $y.key \geq x.key$.
  - 就是说
    - root的值总是大于等于左子树中任意Node值
    - root的值总是小于等于右子树中任意Node值
   
![a](https://imgur.com/Wmi5V8H.jpg)
### by the definition, we have
1.  通过中序遍历(inorder traverse/tree walk), 可以得到由小到大排序过的序列. time: O(n)
2.  对于 BST 的每一个节点 node，它的左侧子树和右侧子树都是 BST (也就是递归的)

### BST 的递归结构

![](https://imgur.com/NlbmByX.jpg)

## operations（全是O(h)）:
### 使用 
- searching (O(h), h是树高): 在树中查找一个值, 返回node指针.
- argmin/ argmax (O(h), h是树高): 返回argmin和argmax的node 
- Successor/ predecessor (O(h), h是树高): input任意node, 返回他在中序遍历返回的sorted list中的下一个、上一个node
### 维护
- insertion (O(h), h是树高): 插入一个node, 维持BST性质
- deletion (O(h), h是树高): 删除一个node, 维持BST性质

## algorithms:
### search()
- 有点像2分搜索但是数据结构已经帮你分好了
```python
def BSTsearch(root, val):
  if root is None or x.key == val:
    return x
  if  x.key > val:
    return BSTsearch(root.left, val)
  if  x.key < val:
    return BSTsearch(root.right, val)
```
### argmin/ argmax: 
- 找到最靠左或最靠右的Node
```python
def BSTargmin(root):
  while root.left is not None:
    root = root.left
  return root

def BSTargmax(root):
  while root.right is not None:
    root = root.right
  return root
```
### Successor/ predecessor
- Successor: 
  - case1: 如果x的右子树非空, x的succssor是它右子树中最左端Node, 即BSTargmin(x.right)
  - case2: 如果x的右子树空, x有可能是某一个左子树中的最右端的Node。 所以我们要向上找到 $y$
    - $y$ is the lowest ancestor of $x$ whose left child is also an ancestor of $x$
    ![](https://imgur.com/vAgXE2T.jpg)
    - 例如上图中13, 我们要找的是15： left child 6 is also ancestor of 13
- predecessor是对称操作
```py
def BSTSuccessor(x):
  if x.right is not None:
    return BSTargmin(x.right)
  else:
    y = x.parent
    while (y is not None) and (x is y.right):
      # should break when x is not y.right, but y.left
      x = y
      y = y.parent
    return y
```
```py
def BSTPredecessor(x):
  if x.left is not None:
    return BSTargmax(x.left)
  else:
    y = x.parent
    while (y is not None) and (x is y.left):
      # should break when x is not y.left, but y.right
      x = y
      y = y.parent
    return y
```
- 因为只需要在树上向上或向下走并且最大的高度不会超过h, 所以复杂度 O(h)

### BSTinsertion(root, newNode)
- 把新的值为newNode.val 的 newNode 插入到合适位置，维持BST性质：
```python
def BSTinsertion(root, newNode):
    # empty tree
    if root is None:  
        root = newNode
        return
    y = None
    x = root
    while x is not None:
        xLast = x  # 结束时x is None, xLast是x变为None前最后的位置
        if newNode.key < x.key:
            x = x.left
        if newNode.key >= x.key:
            x = x.right
    # 此时 xLast 是 newNode 的parent的位置，newNode应该放在x变为None的位置
        newNode.parent = xLast
        if newNode.key < xLast.key:
            xLast.left = newNode
        if newNode.key >= xLast.key:
            xLast.right = newNode 
    return
```
### deletion()
- !!!!!!!递归解在 leetcode 450
- 删除时，删除的点delNode可能有3种情况，分别有3种不同的对应方法
  1. delNode是叶子，没有子树(left = right = None)
     - 直接删除即可
  2. delNode有单个子树
     - 删除掉以后用子树替代掉它的位置就可以 
     - (更改delNode.parent 的指针，更改delNode.child的指针)
  3. delNode有两个子树
     - 在delNode的右子树中找到它的successor(left-most node)。 successor必定只有right没有left子树
     - 如果successor == delNode.right ((c) case)
       - 删除 delNode, 用delNode.right/successor 代替 delNode 的位置
     - 如果successor != delNode.right ((d) case)
       - 删除 delNode, 用successor 代替  delNode 的位置
       - 用successor.right 代替 successor 原来的位置
 
![](https://imgur.com/1aCiGRQ.jpg)

- 使用一个辅助函数 transplant(delNode, replaceNode):
  - 把replaceNode作为树根的任意树代替delNode位置
    - 修改delNode.parent关系
    - 修改replaceNode.parent关系
    - 不修改delNode.children
```python
def transplant(delNode, replaceNode):
    if (delNode.parent is None): 
        return 
    if(delNode is delNode.parent.left):
        delNode.parent.left = replaceNode
    elif(delNode is delNode.parent.right):
        delNode.parent.right = replaceNode

    if (replaceNode is not None):
        replaceNode.parent = delNode.parent
```
```python
def BSTdeletion(delNode):
    # cover first two cases 
    if (delNode.left is None):
        transplant(delNode, delNode.right)
    elif (delNode.right is None):
        transplant(delNode, delNode.left)
    # third case
    else:
        successor = BSTsuccessor(delNode)
        if (successor is delNode.right):
            transplant(delNode, successor)
            # deal with delNode children: only left
            successor.left = delNode.left
            delNode.left.parent = successor
        else:
            transplant(successor, successor.right)
            transplant(delNode, successor)
            # deal with delNode children: left and right
            successor.left = delNode.left
            delNode.left.parent = successor
            successor.right = delNode.right
            delNode.right.parent = successor
```

## final note:
Randomly built binary search tree does not  guarantee to be balanced (ie: height = log(n)). In the worst case the height of the tree is n. To achieve guaranteed balanced tree there are data structures like
- AVL tree: first balanced binary
- B tree: ideal for disk storage、数据索引和数据库索引。 (avoid highly expensive IO operation)
- Red-Black Tree: 广泛运用RAM solution，例如c++ STL map/
- Trie Tree(字典树): 用在统计和排序大量字符串，如自动机, 搜索关键词提示
# Red-Black Tree (RBT):
红黑树是改良的binary search tree，
- 它引入了更多的数据结构规则和更复杂的插入/删除算法
- 但是它保证了近似搜索树平衡性: aka height $h == logn$
- 它保持了插入/删除/查找的的时间效率 O(logn)

与AVL树的不同：
- AVL树是hard constraint： 任意节点的左右子树的高度差的绝对值不大于1。 hard constraint导致基本每次进行插入/删除节点都会破坏平衡性质，导致插入/删除平均性能较差；
- 红黑树是一种大致平衡的二叉搜索树(loose constraint)（叶子到根的最大长度不大于其最小长度之两倍。因此 插入、删除等操作相对不频繁需要fix. ==> 比AVL树更好的write性能。 代价是相对较弱的查找速度

红黑树结点数与树高的约束：lg(n+1) >= h/2

应用：例如 C++11 STL里 set, map，multiset、multimap的实现
### rules of red-black tree: **red-black properties**
1. Every node is either red or black
2. root and NIL-leaves of the tree are black
3. if a node is red, both of its children is black
4. from arbitrary node, the path from itself to any NIL-leaves contain the same number of black nodes.
   - the number is refered as blackHeight(node)
  
- 相比BST, 每个node多用一个field(binary)记录red/black
### operations
 - 搜索的功能与BST完全一致，但是insertion和deletion为了维持红黑性质，与BST不同
- search: O(logn)
- insertion: O(logn)
- deletion: O(logn)

## algorithms
### rotation:
- rotation goal: decrease the height of the tree
  - by move large subtrees up, move small subtrees down
  - rotation should not alter the BST property: small left tree and large right tree
- two types of rotations: left-rotation; right-rotation
  - 2个的算法是对称的
  - time: O(1), 只有修改指针的操作

```python
def leftRotation(x):
    # x.right.left becomes x.right
    newXright = x.right.left  
    x.right = newXright
    if newXright is not None:
        newXright.parent = x

    # x.right becomes the root, and x is its left child
    newRoot = x.right
    newRoot.parent = x.parent
    if (x == x.parent.left):
        x.parent.left = newRoot
    elif (x == x.parent.right)
        x.parent.right = newRoot
    newRoot.left = x
    x.parent = newRoot
```
![](https://imgur.com/CtTEm7J.jpg)
![](https://imgur.com/nsHMGKJ.jpg)

### insertion: O(logn)
- insertion must preserve BST and red-black properties
- strategy
  - insert newNode (according to BSTinsertion) and color it red
  - recolor and rotate nodes to fix violations
    - two possible violations:
    >
        1. root and NIL-leaves of the tree are black
        2. if a node is red, both of its children is black
  - move the pointer up to the next violated tree, and fix new violation. iterate until all violations are fixed
    - this is done by the main iteration, it like the following code
```python
# O(logn)
def RBTinsertion(root, newNode):
BSTinsertion(root, newNode)
newNode.color = RED
# fix part
curNode = newNode
while (curNode.parent is not None) and (curNode.parent.color is RED): # curNode is not root yet; AND there is a violation
    if Cases2AndFixes(): # uncle red case
        curNode = curNode.p.p # make curNode its grandparent after fixing case2
    elif Cases3AndFixes():  # triangle case
        curNode = curNode.parent # the original parent
        fixCases4AndFixes():  # case4 must occur in curNode's original parent after fixing case3
    elif Cases4AndFixes() # line caseL just fix it
    if curNode.parent is None:
        break # escape when reach the root of tree
tree.root.color = BLACK # the case1
```
- 4 only violation scenarios after: insert newNode and color it red
  1. newNode is root 
  2. newNode.uncle is red 
  3. newNode.uncle is black(triangle) 
  4. newNode.uncle is black(line)

### 1. newNode is root 
   - flip the color to black

### 2. newNode.uncle is red 
  - flip z.uncle, z.parent and z.parent.parent
 
|before|after|
|-|-|
|![](https://imgur.com/1cRwd1y.jpg)|![](https://imgur.com/t3iexeb.jpg)|
    
### 3. newNode.uncle is black(triangle)
- newNode, newNode.parent, newNode.grandparent form right triangle: 
  - do rightRotation(newNode.parent)
  
|before|after |
|-|-|
|![](https://imgur.com/waaASsA.jpg)|![](https://imgur.com/TpzncF5.jpg)|


- newNode, newNode.parent, newNode.grandparent form left triangle: 
  - do ieftRotation(newNode.parent)

|before|after |
|-|-|
|![](https://imgur.com/br4ikl7.jpg)|Z took the place of c!!!!!  the image is missing this is a just placeholder |

### 4. newNode.uncle is black(line)
- newNode, newNode.parent, newNode.grandparent form right-direction line: 
  - do leftRotation(newNode.grandparent)
  - swap color of original newNode.parent and newNode.grandparent 

|before|after rot|after color|
|-|-|-|
|![](https://imgur.com/uGC0F7u.jpg)|![](https://imgur.com/8ZshrT9.jpg) |![](https://imgur.com/1E5hw7l.jpg) |

- newNode, newNode.parent, newNode.grandparent form left-direction line: 
  - do rightRotation(newNode.grandparent)
  - swap color of original newNode.parent and newNode.grandparent 
  - not image this time
#### case 3 and 4 can be unified into do rotation of (parent/ grandparent) in the **OPPOSITE DIRECTION OF newNode**


### deletion: O(logn)
 -------TODO
```python
def RBTdeletion(delNode):
    delNode_originColor = delNode.color
    delNodeReplacement = BSTdeletion(delNode) 
    # (delNodeReplacement is the successor/ or L,R tree that took delNode's position)
    if (delNode_originColor is BLACK):
        RBTdeletionFix(delNodeReplacement)

```


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

# union-find

