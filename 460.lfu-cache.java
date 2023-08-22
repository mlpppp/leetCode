/*
 * @lc app=leetcode id=460 lang=java
 *
 * [460] LFU Cache
 */

// @lc code=start

import java.util.HashMap;
import java.util.LinkedHashSet;

class Data {
    int frequency;
    int value;

    public Data(int value) {
        this.frequency = 0;
        this.value = value;
    }

    public int getFrequency() {
        return this.frequency;
    }

    public void setFrequency(int frequency) {
        this.frequency = frequency;
    }

    public int getValue() {
        return this.value;
    }

    public void setValue(int value) {
        this.value = value;
    }


}

class LFUCache {
    int minFreq;
    int capacity;
    HashMap<Integer, Data> keyToData;
    HashMap<Integer, LinkedHashSet<Integer>> freqToKeyRecencyArr;

    
    public LFUCache(int capacity) {
        this.minFreq = 0;
        this.capacity = capacity;
        this.keyToData = new HashMap<>();
        this.freqToKeyRecencyArr = new HashMap<>();
    }
    
    public int get(int key) {
        // return -1 if not exist
        // Data = keyToData[key]
        // addDataFreqByOne(key)

    }
    
    public void put(int key, int value) {
        // 1. update if the key already exist
            // 1.1 update
            // 1.2 add addDataFreqByOne
        // 2. insert if the key-value not already exist
            // check if cache reaches limit
            // 2.1 reaches limit, remove LF & LR Data
                // key = freqToKeyRecencyArr[minFreq].getTail
                // removeData(key)


            // 2.2 not reaches limit,
                // addNewData
            
    }

    // TODO
    private Void removeData(int key) {
        // remove key from keyToData
        // if 
    }

    private void addNewData(int key, int value) {
        // create new Data with frequency 1
        // insert to keyToData
        // insert the key into freqToKeyRecencyArr[1], create a new LinkedHashSet & add freqToKeyRecencyArr if needed
        // if minFreq not 1, update it to 1
    }
    private void addDataFreqByOne(int key) {
        // query the Data
        // update frequency in Data
        // move key of Data to the new LinkedHashSet(freq: +1), create new LinkedHashSet & add freqToKeyRecencyArr if needed
        // update minFreq (+1) if current minimal LinkedHashSet empty, free up the LinkedHashSet 
    }
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache obj = new LFUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
// @lc code=end

