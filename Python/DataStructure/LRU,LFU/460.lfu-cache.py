#
# @lc app=leetcode id=460 lang=python3
#
# [460] LFU Cache
#

# @lc code=start
#
# @lc app=leetcode id=146 lang=python3
#
# [146] LRU Cache
#

# @lc code=start

class LHM_Node:
    def __init__(self, key, val, pNext, pPre) -> None:
        self.key = key
        self.val = val
        self.pNext = pNext
        self.pPre = pPre
        self.feq = 1 

# there are multiple of them, for each frequency
class LinkedList:
    def __init__(self) -> None:
        # container of LHM_Nodes grouped by their frequency
        self.pRecent = LHM_Node('recent', -1, None, None)
        self.pTail = LHM_Node('tail', -1, None, None)
        self.pRecent.pNext = self.pTail
        self.pTail.pPre = self.pRecent

class LinkedHashMap:
    def __init__(self, size) -> None:
        # keyToNode serve as container to all nodes, len(keyToNode) is current cache size
        # feqToLinkList serve as container to all LinkedLists, 

        self.size = size
        self.globalMinFeq = 1 # update when get/update(++cause prevLL empty) or insert (cause freq1LL non empty): gaurantee feqToLinkList[globalMinFeq] non empty
        self.feqToLinkList = {} # update when new freq appear (createNewLinkList())
        self.keyToNode = {} # update when del/insert node

    def insert(self, key, val):
        # delete
        if len(self.keyToNode) == self.size:
            # find least frequently used && least recently used
            if self.size == 0:
                return 
            minLL = self.feqToLinkList[self.globalMinFeq]
            delNode = minLL.pTail.pPre
            if (delNode is minLL.pRecent):
                raise ValueError("self.feqToLinkList[self.globalMinFeq] should not be empty")
            # move out of prevLL
            self.moveNodeOut(delNode)
            # possible require to update globalMinFeq here, tho, as globalMinFeq will soon be updated to 1, no worry
            # delete node from cache
            self.keyToNode.pop(delNode.key)


        # ! insert new node to LL freq == 1
        # create node and hashing
        newNode = LHM_Node(key, val, None, None)
        self.keyToNode[key] = newNode
        curLinkedList = None
        if 1 not in self.feqToLinkList:
            # create LinkedList and hashing
            curLinkedList = self.createNewLinkList(1)
        else:
            curLinkedList = self.feqToLinkList[1]
        # reset globalMinFeq
        self.globalMinFeq = 1
        # insert the node to head
        self.insertNodeToHead(newNode, curLinkedList)


    def update(self, key, val):
        curNode = self.keyToNode[key]
        self.frequentPlusOneOperations(curNode)
        curNode.val = val

    def put(self, key, val):
        if key in self.keyToNode:
            self.update(key, val)
        else:
            self.insert(key, val)

    def get(self, key: int) -> int:
        if key in self.keyToNode:
            curNode = self.keyToNode[key]
            self.frequentPlusOneOperations(curNode)
            return curNode.val
        else:
            return -1
           
    # ! utils
    def frequentPlusOneOperations(self, node):
        # ! update frq++, insert node to new LL and (possibly update global_min)
        # move out of prevLL
        self.moveNodeOut(node)

        # possible require to update globalMinFeq
        if node.feq == self.globalMinFeq:
            # check if prevLL is empty
            prevLL = self.feqToLinkList[node.feq] 
            if prevLL.pRecent.pNext is prevLL.pTail: # LL is empty
              
                self.globalMinFeq += 1 
        # insert curNode to new LinkedList
        node.feq += 1
        # may need to create new LinkedList 
        curLinkedList = None
        if node.feq not in self.feqToLinkList:
            curLinkedList = self.createNewLinkList(node.feq)
        else: 
            curLinkedList = self.feqToLinkList[node.feq]
        self.insertNodeToHead(node, curLinkedList)

        
    def createNewLinkList(self, freq):
        newLinkedList = LinkedList()
        self.feqToLinkList[freq] = newLinkedList
        return newLinkedList

    def insertNodeToHead(self, node, linkList):
        # insert the node to head
            # post node
        node.pNext = linkList.pRecent.pNext
        linkList.pRecent.pNext.pPre = node
            # pRecent(head)
        node.pPre = linkList.pRecent
        linkList.pRecent.pNext = node

    def moveNodeOut(self, node):
        node.pPre.pNext = node.pNext
        node.pNext.pPre = node.pPre


class LFUCache:

    def __init__(self, capacity: int):
        self.LHM = LinkedHashMap(size = capacity)
        
    def get(self, key: int) -> int:
        return self.LHM.get(key)

    def put(self, key: int, value: int) -> None:
        self.LHM.put(key, value)
        


# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
# @lc code=end

# 数据按照访问频次进行排序
        
# 如果多个数据拥有相同的访问频次，删除最早插入的那个数据

# first inserted use counter is set to 1. The use counter is incremented either a get or put operation
    # each increment:
        # add freq by one
        # check if freqLinkList empty, and possibly update global minimal
