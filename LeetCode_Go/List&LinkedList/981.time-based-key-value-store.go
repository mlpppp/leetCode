/*
 * @lc app=leetcode id=981 lang=golang
 *
 * [981] Time Based Key-Value Store
 */

// @lc code=start

// Design a time-based key-value data structure: for each key it can store a time-series value (value at different time stamps) and retrieve the keys' timestamp value. Implement APIS: set(String key, String value, int timestamp): set value at timestamp, get(String key, int timestamp): get value at timestamp, if not exist, get last exist timestamp value. All the timestamps timestamp of set are strictly increasing.

// O(logn) hashmap + binary search. use a hashmap of key->[(value, timestamp)], since timestamp in set(key, value, timestamp) function call are strictly increasing, the [(value, timestamp)] array should be strictly increasing regarding timestamp. Each time get is called we binary search for the left bound of the timestamp in [(value, timestamp)].

type TimeMap struct {
	data map[string][]dataPoint
}

type dataPoint struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	timemap := TimeMap{}
	timemap.data = make(map[string][]dataPoint)
	return timemap
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], dataPoint{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	result := ""
	if arr, exists := this.data[key]; exists {
		// query out of the record bound
		if timestamp < arr[0].timestamp {
			return ""
		}
		// binary search for the timestamp
		left, right := 0, len(arr)-1
		for left <= right {
			mid := (left + right) / 2
			if arr[mid].timestamp > timestamp { // search in the left
				right = mid - 1
			}
			if arr[mid].timestamp <= timestamp { // search in the right
				result = arr[mid].value
				left = mid + 1
			}
		}
	}
	return result
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end

