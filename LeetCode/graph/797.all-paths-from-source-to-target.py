#
# @lc app=leetcode id=797 lang=python3
#
# [797] All Paths From Source to Target
#

# @lc code=start
class Solution:
    result = []
    def traverse(self, graph, lastNodeIdx, rootNode, myRoute):
        # rootNode is actually index
        myRoute.append(rootNode)
        if rootNode == lastNodeIdx:
            self.result.append(myRoute.copy()) 
            myRoute.pop(-1)
            return 
        for node in graph[rootNode]:  # nodes in rootNode.neighbors
            self.traverse(graph, lastNodeIdx, node, myRoute)
        myRoute.pop(-1)
               
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        self.result = []
        lastNodeIdx = len(graph)-1
        self.traverse(graph, lastNodeIdx, 0, [])
        return self.result


# @lc code=end

# 这道题的关键是不使用visited, 从0开始对图进行遍历       
# 因为不使用visited， 在对每一个点进行遍历的时候他们都尝试了所有可能的选择
# 可以使用preVisit(), postVisit()思想动态维护list myRoute, 节约.copy()的花费

# with copy()
    def traverse(self, graph, lastNodeIdx, rootNode, myRoute):
        # rootNode is actually index
        myRoute.append(rootNode)
        print(myRoute)
        if rootNode == lastNodeIdx: 
            self.result.append(myRoute) 
            return 
        for node in graph[rootNode]:  # nodes in rootNode.neighbors
            self.traverse(graph, lastNodeIdx, node, myRoute.copy())
