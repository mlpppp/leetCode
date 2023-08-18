#
# @lc app=leetcode id=701 lang=python3
#
# [701] Insert into a Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        if root is None:
            return TreeNode(val)

        insertionParent_P = root
        insertion_P = root

        while (insertion_P is not None):
            insertionParent_P = insertion_P
            if val > insertion_P.val:
                insertion_P = insertion_P.right
            else:
                insertion_P = insertion_P.left
        if val > insertionParent_P.val:
            insertionParent_P.right = TreeNode(val)
        else:
            insertionParent_P.left = TreeNode(val)
        
        return root   
                
        
# @lc code=end

#! can also solve with recursion !!

# TreeNode insertIntoBST(TreeNode root, int val) {

#     if (root == null) return new TreeNode(val);

#     if (root.val < val) 
#         root.right = insertIntoBST(root.right, val);
#     if (root.val > val) 
#         root.left = insertIntoBST(root.left, val);
#     return root;
# }