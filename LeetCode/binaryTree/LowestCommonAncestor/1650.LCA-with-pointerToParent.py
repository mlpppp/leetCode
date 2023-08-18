from mimetypes import init
# 再看最后一道最近公共祖先的题目吧，力扣第 1650 题「二叉树的最近公共祖先 III」，这次输入的二叉树节点比较特殊，包含指向父节点的指针：
# 
# # 
# 给你输入一棵存在于二叉树中的两个节点 p 和 q，请你返回它们的最近公共祖先，函数签名如下：


def lowestCommonAncestor(p, q);

class Node:
    def __init__(self, val) -> None:
        val = val
        left = None
        right = None
        parent = None

# exactly the same as linkList intersection 
# see leetcode 160