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
    // each root has two paths that includes itself: leftMaximumPath, rightMaximumPath. A root must return max(leftMaximumPath, rightMaximumPath)
    // a maximumPathSum that pass a root is: root + left
    int maxSum = -1001;
    public int maxPathSum(TreeNode root) {
        maxPathSumHelper(root);
        return maxSum;
    }
    public int maxPathSumHelper(TreeNode root) {
        // postorder traverse
        // 1. calculate the max sum that pass the root
        // 2. calculate and return maxSingleHeadedPathSum
        if (root == null) {
            return 0;
        }
        // if a subtree's path sum is negative, it won't be included
        int leftSinglePathSum = Math.max(0, maxPathSumHelper(root.left));
        int rightSinglePathSum = Math.max(0, maxPathSumHelper(root.right));
        // postorder
        int rootMaxPathSum = root.val + leftSinglePathSum + rightSinglePathSum;

        if (rootMaxPathSum > maxSum) { // update global max
            maxSum = rootMaxPathSum;
        }

        return Math.max(root.val+leftSinglePathSum, root.val+rightSinglePathSum);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
