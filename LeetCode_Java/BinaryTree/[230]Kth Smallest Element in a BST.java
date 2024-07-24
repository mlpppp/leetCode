package LeetCode_Java.BinaryTree;
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    // O(n): in order traversal
    // O(logn): change the data structure: in each node, store number of nodes in the left/right subtree
    int counter = 0;
    public int kthSmallest(TreeNode root, int k) {
        counter = 0;
        int result = inorderTraverse(root, k);
        return result;
    }


    public int inorderTraverse(TreeNode root, int k) {
        if (root == null) {
            return -1;
        }
        // left
        int left = inorderTraverse(root.left, k);
        if  (left!= -1) {
            return left;
        }
        // inorder
        counter++;
        if (counter == k) {
            return root.val;
        }
        // right
        int right = inorderTraverse(root.right, k);
        if  (right!= -1) {
            return right;
        }
        return -1;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
