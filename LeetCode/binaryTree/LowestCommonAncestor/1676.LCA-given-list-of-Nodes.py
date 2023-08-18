# 依然给你输入一棵不含重复值的二叉树，但这次不是给你输入 p 和 q 两个节点了，而是给你输入一个包含若干节点的列表 nodes（这些节点都存在于二叉树中），让你算这些节点的最近公共祖先。

# 函数签名如下：


# the solution

def lowestCommonAncestor(root, nodes: List[TreeNode]):
    hashNodes = dict.fromkeys(nodes)
    lowestCommonAncestorRecur(root, hashNodes)


def lowestCommonAncestorRecur(root, hashNodes):
    if root is None: return Non        
    elif (root in hashNodes):
        return root
    else:
        leftResult = lowestCommonAncestor(root.left, hashNodes)
        rightResult = lowestCommonAncestor(root.right, hashNodes)
        if (leftResult and rightResult):
            return root
        if ((not leftResult) and (not rightResult)):
                return None
        return leftResult if leftResult else rightResult

# 算法思想：         
    # findLCA(list ver) logic:
    # 对任意node:


def findLCA(root, hashNodes):
    if root is None: return None
    if root in hashNodes:
        case1: root is LCA
            return root as LCA
        case2: root is LCAchild
            return root as folded children

    if root not in hashNodes: 
        # query l, r trees
        leftResult = findLCA(root.left, p, q)
        rightResult = findLCA(root.right, p, q)

        if both leftResult and rightResult are folded children:
            case 1: root is LCA
                return root as LCA
            case 2: root is local LCA
                return root as folded children
            
        if (either of leftResult or rightResult is folded children):
            return the result as folded children

        if (results are both None):
            return None
 