# - unionFind主要是解决图论中「动态连通性」的数据结构

# - API
#   - union(p, q): 将 p 和 q 所在的connectedComponent连接为一个
#   - find(p): 返回p所在的connectedComponent identifier
#   - boolean isConnected(p, q): 判断 p 和 q 是否连通
#   - count() : 返回图中有多少个连通分量


# UnionFind也是处理等价关系的工具

    # 等价关系，也就是说具有如下三个性质：

    # 1、自反性：节点 p 和 p 是连通的。

    # 2、对称性：如果节点 p 和 q 连通，那么 q 和 p 也连通。

    # 3、传递性：如果节点 p 和 q 连通，q 和 r 连通，那么 p 和 r 也连通。


#! complexity analysis
    # 假设每次总是将size较小的cc union 到 size较大的cc

    # ! union加上之后find压缩的cost是 O(|smaller cc|), 因为所有nodes in smaller cc最终会变成ccParent[larger cc]

    # Consider an element x. what is the worst times of change ?
        # ccParents[x]is changed ≤ log2(n) times
        # cause each time the cc x belongs to at least doubled

    # Therefore, the total time for n − 1 Union operations = O(nlog(n))


class unionFind:
    def __init__(self, ccNum) -> None:
        self.ccCount = ccNum 
        self.ccParents = [i for i in range(ccNum)] # each node point to it's parent
        self.ccSize = {i:1 for i in range(ccNum)}
    def find (self, x):
        # O(min(|smaller cc|) for first use after a union), O(1) for subsequent use
        # return ccParent of node x
        # compress the route by making all nodes in cc point to the same node
        if (self.ccParents[x] != x):
            self.ccParents[x] = self.find(self.ccParents[x])
        return  self.ccParents[x]

    def union(self, p, q):
        # O(1)
        # union connect component node p and q belong to 
        ccParentP = self.find(p)
        ccParentQ = self.find(q)
        if ccParentP == ccParentQ: # already same CC
            return 
        else:
            # 把size小的cc接到size大的cc
            # ! union加上之后find压缩的cost是 |cc of P|, 因为所有nodes in cc of P在压缩后变成ccParentQ
            smallerCC = ccParentP if self.ccSize[ccParentP] < self.ccSize[ccParentQ] else ccParentQ
            largerCC = ccParentQ if smallerCC is ccParentP else ccParentP
            
            # manage ccSize
            smallerSize = self.ccSize[smallerCC]
            self.ccSize[largerCC] += smallerSize
            self.ccSize.pop(smallerCC)

            self.ccParents[smallerCC] = largerCC
            self.ccCount -= 1

    def isConnect(self, p, q):
        # determine if p and q belongs to the same CC
        ccParentP = self.find(p)
        ccParentQ = self.find(q)
        return ccParentP == ccParentQ

    def connectComponentCount(self):
        # num of connect Component in graph
        return self.ccCount

    def connectComponentSize(self, p):
        # O(1)
        # the size of the connect Component node p belongs to 
        parent = self.find(p)
        return self.ccSize[parent]

