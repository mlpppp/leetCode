# traverse

# The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e., there is a directed edge from node i to node graph[i][j]).

# eg: graph = [[1,2],[3],[3],[]]

class graphWalk:
    visited = dict()
    prePost = dict()
    clock = 0
    def traverseDFS(self, graph, rootNode):
        self.visited[rootNode] = True
        print(rootNode)
        self.preVisit(rootNode)
        for nodeNeighbor in graph[rootNode]:
            if nodeNeighbor not in self.visited:
                self.traverse(graph, nodeNeighbor)
            else:
                continue
        self.postVisit(rootNode)

    def preVisit(self, rootNode):
        self.prePost[rootNode] = [self.clock, None]
        self.clock += 1

    def postVisit(self, rootNode):
        self.prePost[rootNode][1] = self.clock
        self.clock += 1

    def walk(self, graph, rootNode):  # master function
        self.visited = dict()
        self.prePost = dict()
        self.clock = 0
        self.traverseDFS(graph, rootNode)

    def traverseBFS(self, graph, rootNode):
    


            
