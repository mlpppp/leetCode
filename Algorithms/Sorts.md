# insertion sort (O(n^2))
## type: brute force
- worst case $O(N^2)$
   
>I: A[1..n]

>O: A[1..n] sorted (incr)

```js
for (let i = 1; i < D.length; i++) {
  const key = D[i];
  let j;
  for (j = i - 1; (j >= 0) && (D[j] > key); j--) {
    D[j + 1] = D[j];
  }
  D[j + 1] = key;
}
```

- 将链表划分为2块，前半为有序区后半为无序区。算法的目的是由初始：len==1 增长有序区，每次从无序区添加一个element将它插入到有序区中正确的位置。

```js
insertion sort()
    for: D[i], i == 2 to D.len
        将D[i]插入到有序区中正确的位置()
```
- 将D[i]插入到有序区中正确的位置(), 是通过将D[i]与之前有序区的, 由大到小每一个Element进行比较。如有序区中比较对象较大，将该对象向右移一位。停止条件：D[i]正确的位置找到（aka序区中比较对象较小）。之后将D[i]插入正确的位置
```js
将D[i]插入到有序区中正确的位置()
    for: D[j], j == i-1 to 0
        terminate when: 
            D[j] <= D[i]  
            OR j == 0     
        move value of D[j] to D[j+1]
    set D[j+1] to D[i]
```

# mergeSort()
## type: Divide n Conquer
- divide：分解原 list (size n) 为 2个 n/2 的子 list
  - mid = floor(lo + high), 
  - L1 = [lo, mid], L2 = [mid+1, high]
- conquer:
  - mergeSort(L1)
  - mergeSort(L2)
- merge:
  - merge(L1, L2)
## 算法
```py
# also leetCode [21]
from math import floor
def mergeTwoSortedList(L1, L2):

  # 2个指针分别指向2个数组头， 
  # 每次向output里面插入2个中较小的一个argmin， 并移动argmin的指针
  # 如果其中一个指针遍历完了就把另一个指针剩下的元素全部插入到output中

    p1 = p2 = 0
    output = []
    while (p1 < len(L1) and p2 < len(L2)):
        if (L1[p1] <= L2[p2]):
            output.append(L1[p1])
            p1+=1
        else:
            output.append(L2[p2])
            p2+=1
    if (p1 == len(L1)):
        output.extend(L2[p2:len(L2)])
    elif (p2 == len(L2)):
        output.extend(L1[p1:len(L1)])
    return output

def mergeSort(L):
    # unit case
    if(len(L)<=1): return L

    # divide
    mid = floor((0+len(L)-1)/2)
    L1 = L[0:mid+1]
    L2 = L[mid+1:len(L)]

    # conquer sub problem
    L1 = mergeSort(L1)
    L2 = mergeSort(L2)

    # combine solution
    L = mergeTwoSortedList(L1, L2)
    return L
```
## complexity
- mergeTwoSortedList(): O(n)
- mergeSort(): T(n)
  - T(n) = 2T(n/2) + O(n)  
  - ==> O(nlogn)
- Auxiliary Space: O(n)
- not in-place
# heapSort()
see [here](../DataStructure/AdvStructures.md#Heap)

# QuickSort(), 原址，期望O(nlogn)
### QuickSort(A, left, right)
- input: list A, left, right 定义想要排序的A子数组区间
  - left = 0, right = len(A)-1 就是整个数组
  
### partition(A, left, right) -> mid
- 通过partition方法，在 A 中 [left, right] 区间找到一个归位元素index mid: 使得mid满足
  - 对任意在 A[left, mid-1] 中元素 q, A[q] < A[mid]
  - 对任意在 A[mid+1, right] 中元素 q, A[mid] < A[q] 
- 此时mid应该是在output的sorted数组中相同的位置： 归位

- partition算法
  - A[mid] 使用的值是初始A[right]的值 (pivot value)
  - 双指针算法：pEdit 和 pSearch
    - pSearch寻找比pivot小的元素, 如果找到了就与当前的pEdit交换， 并且pEdit++
    - pEdit 永远都指向pivot小的元素区域的最前端
```py
def partition(A, left, right):
    pivot = A[right]
    pEdit = pSearch = 0
    while (pSearch < right):
      if (A[pSearch] < pivot):
        swap (A[pEdit], A[pSearch])
        pEdit+=1
      pSearch+=1

    swap(A[pEdit], A[right])  # 最后使pivot归位
    return pEdit
```
### QuickSort(): divide and conquer
- divide: 通过 partition(A, left, right) 找到一个归位元素， 并且将A分解为 L1: [left, mid-1] L2: [mid+1, right]
- conquer: QuickSort(L1)和QuickSort(L2)
- combine: 因为是原址排序并不需要combine，每一次调用QuickSort()都会有一个元素归位

```py
def QuickSort(A, left, right):
    if left >= right: return None
    mid = partition(A, left, right)
    QuickSort(A, left, mid-1)
    QuickSort(A, mid+1, right)
QuickSort(A, 0, len(A)-1)
```

### complexity:
最差 O(n^2)
平均 O(nlogn)
空间O(1)