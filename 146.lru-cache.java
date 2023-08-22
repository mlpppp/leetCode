/*
 * @lc app=leetcode id=146 lang=java
 *
 * [146] LRU Cache
 */

// @lc code=start

import java.util.HashMap;

class LinkedListNode {
    int key;
    int value;
    LinkedListNode nextNode;
    LinkedListNode prevNode;

    public LinkedListNode(int key, int value, LinkedListNode nextNode, LinkedListNode prevNode) {
        this.key = key;
        this.value = value;
        this.nextNode = nextNode;
        this.prevNode = prevNode;
    }

    public int getKey() {
        return this.key;
    }

    public void setKey(int key) {
        this.key = key;
    }

    public int getValue() {
        return this.value;
    }

    public void setValue(int value) {
        this.value = value;
    }

    public LinkedListNode getNextNode() {
        return this.nextNode;
    }

    public void setNextNode(LinkedListNode nextNode) {
        this.nextNode = nextNode;
    }

    public LinkedListNode getPrevNode() {
        return this.prevNode;
    }

    public void setPrevNode(LinkedListNode prevNode) {
        this.prevNode = prevNode;
    }
}

class LinkedHashMap {
    int size;
    LinkedListNode head;  // most recent, no data
    LinkedListNode tail; // no data
    HashMap<Integer, LinkedListNode> hashmap; 

    public void initLRU(int capacity) {
        size =  capacity;
        hashmap = new HashMap<>();
        head = new LinkedListNode(-1, -1, null, null);        
        tail = new LinkedListNode(-1, -1, null, null);
        head.setNextNode(tail);        
        tail.setPrevNode(head);
    }

    public void put(int key, int value) {
        // 1. update if key exist 
        LinkedListNode node = hashmap.get(key);
        if (node != null) {
            node.setValue(value);
            moveoutNode(node);
            insertNodeAfter(node, this.head);
            return;
        }
        // 2. add to cache if key not exist
        // 2.1 cache size exceed: evict
        if (hashmap.size() >= this.size) {
            LinkedListNode delNode = this.tail.prevNode;
            moveoutNode(delNode);
            hashmap.remove(delNode.getKey());
            delNode = null; // remove all ref to the node
        }
        // 2.2 cache size not exceed: insert to head
        LinkedListNode newNode = new LinkedListNode(key, value, null, null);
        // new node will be put to the head
        insertNodeAfter(newNode, this.head);
        hashmap.put(key, newNode);
    }
    public int get(int key) {
        LinkedListNode node = hashmap.get(key);
        if (node == null) {
            return -1;
        }
        // move node to the head
        moveoutNode(node);
        insertNodeAfter(node, this.head);
        return node.getValue();
    }

    // linkedList utils
    private void insertNodeAfter(LinkedListNode insertedNode, LinkedListNode targetNode) {
        targetNode.nextNode.prevNode = insertedNode;
        insertedNode.nextNode = targetNode.nextNode;
        targetNode.nextNode = insertedNode;
        insertedNode.prevNode = targetNode;
    }

    private void moveoutNode(LinkedListNode node) {
        LinkedListNode rightNode = node.nextNode;        
        LinkedListNode leftNode = node.prevNode;
        leftNode.nextNode = rightNode;
        rightNode.prevNode = leftNode;
    }

}


class LRUCache {
    LinkedHashMap linkedHashMap;
    public LRUCache(int capacity) {
        linkedHashMap = new LinkedHashMap();
        linkedHashMap.initLRU(capacity);
    }
    
    public int get(int key) {
        return linkedHashMap.get(key);
    }
    
    public void put(int key, int value) {
        linkedHashMap.put(key, value);
    }
    public static void main(String[] args) {
        LRUCache test = new LRUCache(2);
        test.put(1, 1);
        test.put(2, 2);

        test.get(1);
        System.out.println(test.get(3));
    }
    
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
// @lc code=end

