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
    // assume left and right subtrees are both bst
    // need to make sure rightmost node in left subtree and leftmost node in right subtree enclose the root
    public boolean isValidBST(TreeNode root) {
        if (root == null) return true;

        boolean leftIsGood = false;
        if (root.left == null) {
            leftIsGood = true;
        } else {
            if (isValidBST(root.left)) {
                // find the rightmost node in left subtree
                TreeNode leftCur = root.left;
                while (leftCur.right != null) {
                    leftCur = leftCur.right;
                }
                int leftLargest = leftCur.val;
                if (leftLargest < root.val) {
                    leftIsGood = true;
                }
            }
        }
        // find the leftmost node in right subtree
        boolean rightIsGood = false;
        if (root.right == null) {
            rightIsGood = true;
        } else {
            if (isValidBST(root.right)) {
                TreeNode rightCur = root.right;
                while (rightCur.left != null) {
                    rightCur = rightCur.left;
                }
                int rightSmallest = rightCur.val;
                if (rightSmallest > root.val) {
                    rightIsGood = true;
                }
            }
        }
        // BST criterion
        return leftIsGood && rightIsGood;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
