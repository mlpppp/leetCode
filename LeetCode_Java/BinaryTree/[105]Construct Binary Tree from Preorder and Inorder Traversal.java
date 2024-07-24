package LeetCode_Java.BinaryTree;
//leetcode submit region begin(Prohibit modification and deletion)

import java.util.HashMap;

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
//    HashMap<Integer, Integer> preorderIdx = new HashMap<>();
    HashMap<Integer, Integer> inorderValToIdx = new HashMap();
    public TreeNode buildSubTree(int[] preorder, int preStart, int preEnd, int[] inorder, int inStart, int inEnd) {

        // base case:
        if (preStart == preEnd || inStart == inEnd) {
            return null;
        }

        // find root in preorder and inorder
        int rootVal = preorder[preStart];
        TreeNode root = new TreeNode(rootVal);;
        int rootIdx = inorderValToIdx.get(rootVal);

        // calculate left/right subtrees from root location
        int lInStart = inStart;
        int lInEnd = rootIdx;
        int lTreeSize = rootIdx - inStart;
        int lPreStart = preStart + 1;
        int lPreEnd = preStart + 1 + lTreeSize;
        int rInStart = rootIdx + 1;
        int rInEnd = inEnd;
        int rPreStart = preStart + lTreeSize + 1;
        int rPreEnd = preEnd;

        // traverse to build tree
        root.left =  buildSubTree(preorder, lPreStart, lPreEnd, inorder, lInStart, lInEnd);
        root.right =  buildSubTree(preorder, rPreStart, rPreEnd, inorder, rInStart, rInEnd);

        return root;
    }
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        // use hashmap to store the index of node values for quick access
        for (int i = 0; i < preorder.length; i++) {
            inorderValToIdx.put(inorder[i], i);
        }

        return buildSubTree(preorder, 0, preorder.length, inorder, 0, inorder.length);
    }
}
//leetcode submit region end(Prohibit modification and deletion)
