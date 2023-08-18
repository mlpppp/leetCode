#
# @lc app=leetcode id=785 lang=python3
#
# [785] Is Graph Bipartite?
#

# 判定二分图的算法很简单，就是用代码解决「双色问题」
    # 说白了就是遍历一遍图，一边遍历一边染色（两个颜色反复换），看看能不能用两种颜色给所有节点染色，且相邻节点的颜色都不相同。

    # 关键是判断相邻的，visited node， 如果颜色不同，马上可以返回false

# @lc code=start
class Solution:
    isBipartite = True
    def solveBiColor_DFS(self, graph, rootNodeIdx, visited, colors):
        # preVisit(): escape if already know isnotBipartite
        if not self.isBipartite:
            return  

        visited[rootNodeIdx] = True

        # visit
        for nextNodeIdx in graph[rootNodeIdx]:
            if not visited[nextNodeIdx]:
                colors[nextNodeIdx] = not colors[rootNodeIdx]  # set next color
                self.solveBiColor_DFS(graph, nextNodeIdx, visited, colors)
            else:
                # the node has been visited
                if (colors[nextNodeIdx] == colors[rootNodeIdx]):
                    self.isBipartite = False
                    return 
          
        #postVisit()
    def DFStraverse(self, graph, rootNodeIdx, visited):
        # preVisit()
        visited[rootNodeIdx] = True
        # visit
        for nextNodeIdx in graph[rootNodeIdx]:
            if not visited[nextNodeIdx]:
                self.DFStraverse(graph, nextNodeIdx, visited)

    def findConnectedComponent_undir(self, graph):
        # return list of node(number == numofConnectedComponent), each belongs to is ConnectedComponent
        visited = [False for i in range(len(graph))]
        ConnectedComponent = []
        for nodeIdx in range(len(graph)):
            if not visited[nodeIdx]:
                ConnectedComponent.append(nodeIdx)
                self.DFStraverse(graph, nodeIdx, visited)
        return ConnectedComponent

          


        
    def isBipartite(self, graph: List[List[int]]) -> bool:
        self.isBipartite = True
        initNodesIdx = self.findConnectedComponent_undir(graph)
        for nodeIdx in initNodesIdx:
            visited = [False for i in range(len(graph))]
            colors = [False for i in range(len(graph))]
            self.solveBiColor_DFS(graph, nodeIdx, visited, colors)

        return self.isBipartite


# @lc code=end

