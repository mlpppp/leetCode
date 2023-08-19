#
# @lc app=leetcode id=210 lang=python3
#
# [210] Course Schedule II
#

# @lc code=start
class graphNode:
    def __init__(self, val, postNodes, preNodes) -> None:
        self.val = val
        self.postNodes = postNodes
        self.preNodes = preNodes

        self.preNumber = None
        self.postNumber = None

        self.disabled = False  # Topology order util
    def initPrePost(self):
        self.preNumber = None
        self.postNumber = None

class Solution:
    visited = {}
    clock = 0
    def toGraph(self, numCourses, prerequisites):
        #! return: list of graphNode(class), with order (0 to n-1)
        graph = [graphNode(val=i, postNodes=[], preNodes=[]) for i in range(numCourses)]
        for edge in prerequisites:
            fromNode = edge[1]
            toNode = edge[0]
            graph[fromNode].postNodes.append(toNode)
            graph[toNode].preNodes.append(fromNode)
        return graph

    # traverse utils
    def preVisit(self, rootNode:graphNode):
        rootNode.preNumber = self.clock
        self.clock+=1
        
    def postVisit(self, rootNode:graphNode):
        rootNode.postNumber = self.clock
        self.clock+=1
        

    def traverse(self, graph, rootNodeIdx):
        #! traverse graph from node: rootNodeIdx
        rootNode = graph[rootNodeIdx]  # graphNode
        # preVisit
        self.preVisit(rootNode)  

        # visit
        self.visited[rootNodeIdx] = True
        
        # recur
        for nodeIdx in rootNode.postNodes: #:graphNode
            if nodeIdx not in self.visited:
                self.traverse(graph, nodeIdx)      
            else:
                continue
        
        # postVisit
        self.postVisit(rootNode)

    def traverseInit(self, graph):
        for node in graph:
            node.initPrePost()
        self.visited = {}
        self.clock = 0

    def isAcyclic(self, graph, prerequisites):
        for firstNodeIdx in range(len(graph)):
            # initialize traverse
            self.traverseInit(graph)          
            # traverse
            self.traverse(graph, rootNodeIdx=firstNodeIdx)
 
            # test cyclic        
            for edge in prerequisites:
                fromNode = graph[edge[1]]  # type: graphNode
                toNode = graph[edge[0]] # type: graphNode
                if (fromNode.preNumber is not None and toNode.preNumber is not None):
                    if (fromNode.preNumber >= toNode.preNumber) and (fromNode.postNumber <= toNode.postNumber) :
                        return False
        return True

    def canFinish(self, graph, prerequisites) -> bool:
        return self.isAcyclic(graph, prerequisites)

#!-----------------------findOrder utils---------------------------

    def findZeroIndegreeNodes(self, graph):
        #! return list of indices
        result = []
        for node in graph:
            if (len(node.preNodes) == 0) and (not node.disabled):
                result.append(node.val)
        return result

    def deleteNode(self, graph, delIdx):
        #! delete node:delIdx, and also edges from it
        delNode = graph[delIdx]
        # remove delNode from node.preNodes where node in delNode.postNodes
        for nodeIdx in delNode.postNodes:
            curNode = graph[nodeIdx]
            curNode.preNodes.remove(delIdx)
        delNode.disabled = True
        

    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        graph = self.toGraph(numCourses, prerequisites)
        if not self.canFinish(graph, prerequisites):
            return []
        # topology order
        result = []
        while (len(result) < len(graph)): # len(result) ++, should terminate
            zeroNodes =  self.findZeroIndegreeNodes(graph)
            curNodeIdx = zeroNodes.pop(0)
            result.append(curNodeIdx)
            self.deleteNode(graph, curNodeIdx)  
        return result


# @lc code=end

# sol1: 
# 当结果的长度不是所有node的时候
    #  找到一个list，零入度的点，把第一个加到结果
    #  接着把这一个点删除并且把他的所有相关的边删除把边删除就需要把所有node的pre node中这一个边删掉

# sol2 (theoretical):
    # Step 1 : Run DFS (including pre/post-numbers)
    # Step 2 : Run Bucket Sort to sort the vertices by postnumber.
    # Step 3 : Obtain the topological ordering from the reverse sorted
    # order of the postnumbers