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
    boolean exist = false;
    public boolean isSubtree(TreeNode root, TreeNode subRoot) {
        if (exist) {
            return true;
        }
        if (root == null) return false;
        if (isSameTree(root, subRoot) ||
                isSubtree(root.left, subRoot) ||
                isSubtree(root.right, subRoot)) {
            exist = true;
            return true;
        }
        return isSameTree(root, subRoot);
    }



    public boolean isSameTree(TreeNode p, TreeNode q) {
        // null case
        if (p == null || q == null) {
            return p == q;
        }
        // the root is the same
        if (p.val != q.val) return false;
        // and left/ right subtree are the same tree
        return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
