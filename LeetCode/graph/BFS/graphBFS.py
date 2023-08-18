

def BFS(startNode, visited):
    toVisit = [node for node in startNode.adjNodes]
    # toVisitOrVisited: also add all nodes in toVisit to avoid repeatly add node that have not yet been visited but already in toVisit
    toVisitOrVisited = {startNode:True}  
    for node in toVisit:
        toVisitOrVisited[node] = True

    print(startNode)
    while len(toVisit) > 0:
        curNode = toVisit.pop(0)
        print(curNode)

        for node in curNode.adjNodes:
            if node not in toVisitOrVisited:
                toVisit.append(node) 
                toVisitOrVisited[node] = True


def shortestPath_BFS(startNode, targetNode):
    # return min(#steps) from startNode to targetNode
    if startNode is targetNode:
        return 0

    toVisit = [node for node in startNode.adjNodes]
    toVisitOrVisited = {startNode:True}  
    for node in toVisit:
        toVisitOrVisited[node] = True

    step = 1
    while len(toVisit) > 0:
        xStepsReachable = len(toVisit)  # an extra layer of loop to get the set of nodes reachable within x steps from startNode
        for nodeIdx in range(xStepsReachable):
            curNode = toVisit.pop(0)
            if curNode is targetNode:
                return step
            else:
                for node in curNode.adjNodes:
                    if node not in toVisitOrVisited:
                        toVisit.append(node) 
                        toVisitOrVisited[node] = True
        step+=1
