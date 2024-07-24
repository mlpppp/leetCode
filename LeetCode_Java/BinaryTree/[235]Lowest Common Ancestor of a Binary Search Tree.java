package LeetCode_Java.BinaryTree;
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */

class Solution {
    // O(logn). make sure p < q
    // three possiblilities:
        // p < root && q > root, return root
        // p or q is the root, return root
        // p, q both < or > root, find lca in the subtree
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (p.val > q.val) {  // swap p and q reference if p > q
            TreeNode tmp = p;
            p = q;
            q = tmp;
        }
        if (root.val < p.val) {
            return lowestCommonAncestor(root.right, p, q);
        } else if (root.val > q.val){
            return lowestCommonAncestor(root.left, p, q);
        } else {
            return root;
        }
    }


}
//leetcode submit region end(Prohibit modification and deletion)
