#
# @lc app=leetcode id=215 lang=python3
#
# [215] Kth Largest Element in an Array
#

# @lc code=start
from random import shuffle
class Solution:
    def swap(self, nums, i, j):
        if i == j:
            return 
        else: 
            temp = nums[i]
            nums[i] = nums[j]
            nums[j] = temp
            return
    def partition(self, nums, left, right):
        # O(n)
        pivot = nums[right]
        pEdit = pSearch= left
        while (pSearch < right):
            if (nums[pSearch] <= pivot):
                self.swap(nums, pEdit, pSearch)
                pEdit += 1
            pSearch += 1
        self.swap(nums, pEdit, right)
        return pEdit

    def findKthLargestRecur(self, nums: List[int], k: int, left, right) -> int:
        myRank = self.partition(nums, left, right)
        if myRank == len(nums)-k:
            return myRank
        if myRank > len(nums)-k:
            return self.findKthLargestRecur(nums, k, left, myRank-1)
        elif myRank < len(nums)-k:     
            return self.findKthLargestRecur(nums, k, myRank+1, right)

    def findKthLargest(self, nums: List[int], k: int) -> int:
        shuffle(nums)
        idx = self.findKthLargestRecur(nums, k, 0, len(nums)-1)
        return nums[idx]


# 用了和快速排序相同的partition算法每一次partition可以得到一个绝对的归位的值他的位置就是他的rank（nth smallest)， 这样我们就可以在被划分的区间里面递归的找想要的findKthLargest

# best case
    # T(n) = T(n/2) + n
    # N + N/2 + N/4 + N/8 + ... + 1 = 2N = O(N)
# worst case
    # N + (N - 1) + (N - 2) + ... + 1 = O(N^2)
            
# @lc code=end

# heap sol
import heapq

# def findKthLargest(self, nums: List[int], k: int) -> int:
#         heap = []
#         for num in nums:
#             heapq.heappush(heap, num)
#             if (len(heap)>k):
#                 heapq.heappop(heap)
#         return heapq.heappop(heap)

# 总的时间复杂度就是 O(Nlogk) n次insert
# 空间复杂度很显然就是二叉堆的大小，为 O(k)。