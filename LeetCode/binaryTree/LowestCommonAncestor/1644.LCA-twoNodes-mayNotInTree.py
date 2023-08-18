# 1644 题「二叉树的最近公共祖先 II」：

# 给你输入一棵不含重复值的二叉树的，以及两个节点 p 和 q。P和Q有可能不存在于树中！！！！如果 p 或 q 不存在于树中，则返回空指针，否则的话返回 p 和 q 的最近公共祖先节点。
# 函数签名如下：
# solution
class Solution:
    # def lowestCommonAncestorRecur(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode'):
    pExist = False
    qExist = False

    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        result = lowestCommonAncestorRecur(root, p, q)
        if self.pExist and self.qExist:
            return result
        else: 
            return None

    def lowestCommonAncestorRecur(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if root is None: return None

        # 先不管reversion返回的是什么
        leftResult = self.lowestCommonAncestorRecur(root.left, p, q)
        rightResult = self.lowestCommonAncestorRecur(root.right, p, q)

        if leftResult and rightResult:
            # root is LCA, pExist/qExist taken care in chidren
            return root

        if root is p or root is q:
            if root is p:
                self.pExist=True 
            if root is q:
                self.qExist=True 
            return root


        if ((not leftResult) and (not rightResult)):
            return None
        return leftResult if leftResult else rightResult 
