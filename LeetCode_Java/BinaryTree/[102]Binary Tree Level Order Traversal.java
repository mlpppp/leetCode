package LeetCode_Java.BinaryTree;
//leetcode submit region begin(Prohibit modification and deletion)

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

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
    // 102 BFS, level order traverse binary tree

    public List<List<Integer>> levelOrder(TreeNode root) {
        // initialize
        List<List<Integer>> result = new ArrayList<>();
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        if (root != null) {
            queue.add(root);  // careful null can still be count as a queue element
        }
        int level = 0;
        while(queue.size() > 0) {
            level++;
            Queue<TreeNode> nextQueue = new LinkedList<TreeNode>();
            result.add(new ArrayList<>());
            // iterate current level
            while(queue.size() > 0) {
                TreeNode node = queue.remove();
                result.get(result.size() - 1).add(node.val);
                if (node.left != null) {
                    nextQueue.add(node.left);
                }
                if (node.right!= null) {
                    nextQueue.add(node.right);
                }
            }
            queue = nextQueue;
        }
        return result;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
