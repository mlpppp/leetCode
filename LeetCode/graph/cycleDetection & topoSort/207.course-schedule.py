#
# @lc app=leetcode id=207 lang=python3
#
# [207] Course Schedule
#

# @lc code=start

class graphNode:
    def __init__(self, val, neighbors) -> None:
        self.val = val
        self.neighbors = neighbors
        self.preNumber = None
        self.postNumber = None
    def initPrePost(self):
        self.preNumber = None
        self.postNumber = None


class Solution:
    visited = {}
    clock = 0
    def toGraph(self, numCourses, prerequisites):
        graph = [graphNode(val=i, neighbors=[]) for i in range(numCourses)]
        for edge in prerequisites:
            graph[edge[1]].neighbors.append(edge[0])
        return graph

    # traverse utils
    def preVisit(self, rootNode:graphNode):
        rootNode.preNumber = self.clock
        self.clock+=1
        
    def postVisit(self, rootNode:graphNode):
        rootNode.postNumber = self.clock
        self.clock+=1
        

    def traverse(self, graph, rootNodeIdx):
        rootNode = graph[rootNodeIdx]  # graphNode
        # preVisit
        self.preVisit(rootNode)  

        # visit
        self.visited[rootNodeIdx] = True
        
        # recur
        for nodeIdx in rootNode.neighbors: #:graphNode
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

 

    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph = self.toGraph(numCourses, prerequisites)
        return self.isAcyclic(graph, prerequisites)


        
# @lc code=end




# https://labuladong.github.io/algo/2/22/51/

// 记录一次递归堆栈中的节点
boolean[] onPath;
// 记录遍历过的节点，防止走回头路
boolean[] visited;
// 记录图中是否有环
boolean hasCycle = false;

boolean canFinish(int numCourses, int[][] prerequisites) {
    List<Integer>[] graph = buildGraph(numCourses, prerequisites);
    
    visited = new boolean[numCourses];
    onPath = new boolean[numCourses];
    
    for (int i = 0; i < numCourses; i++) {
        // 遍历图中的所有节点
        traverse(graph, i);
    }
    // 只要没有循环依赖可以完成所有课程
    return !hasCycle;
}

void traverse(List<Integer>[] graph, int s) {
    if (onPath[s]) {
        // 出现环
        hasCycle = true;
    }
    
    if (visited[s] || hasCycle) {
        // 如果已经找到了环，也不用再遍历了
        return;
    }
    // 前序代码位置
    visited[s] = true;
    onPath[s] = true;
    for (int t : graph[s]) {
        traverse(graph, t);
    }
    // 后序代码位置
    onPath[s] = false;
}

List<Integer>[] buildGraph(int numCourses, int[][] prerequisites) {
    // 代码见前文
}