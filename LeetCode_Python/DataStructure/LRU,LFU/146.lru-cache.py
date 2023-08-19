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
        
class LinkedHashMap:
    def __init__(self, size) -> None:
        self.size = size
        self.hashMap = {}
        self.pRecent = LHM_Node('recent', -1, None, None)
        self.pTail = LHM_Node('tail', -1, None, None)
        self.pRecent.pNext = self.pTail
        self.pTail.pPre = self.pRecent

    def get(self, key):
        if key in self.hashMap:
            node = self.hashMap[key] 
            # !move node to most recent 
            # move out node
            node.pPre.pNext = node.pNext
            node.pNext.pPre = node.pPre
            # move node in front
            self.pRecent.pNext.pPre = node
            node.pNext = self.pRecent.pNext
            self.pRecent.pNext = node
            node.pPre = self.pRecent

            return node.val
        else:
            return -1

    def insert(self, key, val):
            if len(self.hashMap) == self.size:
                lastNode = self.pTail.pPre
                # manage pointers
                lastNode.pPre.pNext = self.pTail
                self.pTail.pPre = lastNode.pPre
                # del hash and node
                self.hashMap.pop(lastNode.key)
                del lastNode

            
            newNode = LHM_Node(key, val, None, None)
            # manage pointers
            newNode.pNext = self.pRecent.pNext
            self.pRecent.pNext.pPre = newNode

            self.pRecent.pNext = newNode
            newNode.pPre = self.pRecent

            # create Hash
            self.hashMap[key] = newNode
            

        

    def update(self, key, val):
        node = self.hashMap[key] 
        node.val = val

        # !move node to most recent 
        # move out node
        node.pPre.pNext = node.pNext
        node.pNext.pPre = node.pPre
        # move node in front
        self.pRecent.pNext.pPre = node
        node.pNext = self.pRecent.pNext
        self.pRecent.pNext = node
        node.pPre = self.pRecent

    def put(self, key, val):
        
        if key in self.hashMap:
            self.update(key, val)
 
        else:
            self.insert(key, val)

            

           

        

class LRUCache:

    def __init__(self, capacity: int):
        self.LHM = LinkedHashMap(size = capacity)
        
        

    def get(self, key: int) -> int:
        return self.LHM.get(key)
        

    def put(self, key: int, value: int) -> None:
        self.LHM.put(key, value)

        


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
# @lc code=end

# https://labuladong.github.io/algo/2/23/59/

# O(1) I/O， 能用hashtable解决
# 为了让hashtable具有时序性，结合linklist为hashtable添加时序性
# ==> LinkedHashMap