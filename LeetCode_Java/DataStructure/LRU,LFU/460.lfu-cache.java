/*
 * @lc app=leetcode id=460 lang=java
 *
 * [460] LFU Cache
 */

// @lc code=start

import java.util.HashMap;
import java.util.LinkedHashSet;

class Data {
    private int frequency;
    private int value;

    public Data(int value) {
        this.frequency = 1;
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
        if (capacity<=0) {
            System.err.println("Invaild capacity"); 
        }
        this.minFreq = 0;
        this.capacity = capacity;
        this.keyToData = new HashMap<>();
        this.freqToKeyRecencyArr = new HashMap<>();
    }


    
    public int get(int key) {
        // return -1 if not exist
        if (!this.keyToData.containsKey(key)) {
            return -1;
        }
        // Data = keyToData[key]
        Data data = keyToData.get(key);
        // addDataFreqByOne(key)
        addDataFreqByOne(key);
        return data.getValue();
    }
    
    public void put(int key, int value) {
        // 1. update if the key already exist
        if (this.keyToData.containsKey(key)) {
            Data updateData = keyToData.get(key);
            updateData.setValue(value);     // update
            this.addDataFreqByOne(key);     // add addDataFreqByOne
            return;
        }

        // 2. insert if the key-value not already exist
        // check if cache reaches limit
        boolean isReachingLimit = (keyToData.size() >= this.capacity);
        // 2.1 reaches limit, remove LF & LR Data
        if (isReachingLimit) {
            // key = freqToKeyRecencyArr[minFreq].getLeastRecency
            LinkedHashSet<Integer> minFreqKeySet = this.freqToKeyRecencyArr.get(this.minFreq);
            int invalidateKey = minFreqKeySet.iterator().next(); 
            // removeKey(key)
            removeKey(invalidateKey);
            addNewData(key, value);
        }
        // 2.2 not reaches limit,
        else {
            // addNewData
            addNewData(key, value);
        }
    }

    private void removeKey(int key) {
        // remove key from freqToKeyRecencyArr
        int keyFreq = keyToData.get(key).getFrequency();
        this.freqToKeyRecencyArr.get(keyFreq).remove(key);
        // remove key from keyToData
        this.keyToData.remove(key);
        // remove data(auto deref after function finished)
    }

    private void addNewData(int key, int value) {
        // create new Data with frequency 1
        Data newData = new Data(value);
        newData.setFrequency(1);
        // insert to keyToData
        keyToData.put(key, newData);
        // insert the key into freqToKeyRecencyArr[1], create a new LinkedHashSet + addTo freqToKeyRecencyArr if freqToKeyRecencyArr[1] not exist
        if (!freqToKeyRecencyArr.containsKey(1)) {
            LinkedHashSet<Integer> newKeySet = new LinkedHashSet<>();
            freqToKeyRecencyArr.put(1, newKeySet);
        }
        freqToKeyRecencyArr.get(1).add(key);
        // update minFreq to 1
        this.minFreq = 1;
    }

    private void addDataFreqByOne(int key) {
        // query the Data
        // update frequency in Data
        Data data = keyToData.get(key); 
        int prevFreq = data.getFrequency();
        int newFreq = prevFreq+1;
        data.setFrequency(newFreq);
        // move Data key to the new LinkedHashSet(freq: +1), create new LinkedHashSet + addTo freqToKeyRecencyArr if needed
        if (!freqToKeyRecencyArr.containsKey(newFreq)) {
            LinkedHashSet<Integer> newKeySet = new LinkedHashSet<>();
            freqToKeyRecencyArr.put(newFreq, newKeySet);
        }
        freqToKeyRecencyArr.get(newFreq).add(key);

        // free up the prev LinkedHashSet if it is empty, 
            // update minFreq (+1) if the prev LinkedHashSet is minimal freq
        freqToKeyRecencyArr.get(prevFreq).remove(key);
        if (freqToKeyRecencyArr.get(prevFreq).isEmpty()) {
            if (this.minFreq == prevFreq) {
                this.minFreq = newFreq;
            }
            freqToKeyRecencyArr.remove(prevFreq);
        }
    }
}


// class TestLFUCache() {
    
// }
/**
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache obj = new LFUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
// @lc code=end

