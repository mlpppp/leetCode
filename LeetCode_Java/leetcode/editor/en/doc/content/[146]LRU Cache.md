<p>Design a data structure that follows the constraints of a <strong><a href="https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU" target="_blank">Least Recently Used (LRU) cache</a></strong>.</p>

<p>Implement the <code>LRUCache</code> class:</p>

<ul> 
 <li><code>LRUCache(int capacity)</code> Initialize the LRU cache with <strong>positive</strong> size <code>capacity</code>.</li> 
 <li><code>int get(int key)</code> Return the value of the <code>key</code> if the key exists, otherwise return <code>-1</code>.</li> 
 <li><code>void put(int key, int value)</code> Update the value of the <code>key</code> if the <code>key</code> exists. Otherwise, add the <code>key-value</code> pair to the cache. If the number of keys exceeds the <code>capacity</code> from this operation, <strong>evict</strong> the least recently used key.</li> 
</ul>

<p>The functions <code>get</code> and <code>put</code> must each run in <code>O(1)</code> average time complexity.</p>

<p>&nbsp;</p> 
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input</strong>
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
<strong>Output</strong>
[null, null, null, 1, null, -1, null, -1, 3, 4]

<strong>Explanation</strong>
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
</pre>

<p>&nbsp;</p> 
<p><strong>Constraints:</strong></p>

<ul> 
 <li><code>1 &lt;= capacity &lt;= 3000</code></li> 
 <li><code>0 &lt;= key &lt;= 10<sup>4</sup></code></li> 
 <li><code>0 &lt;= value &lt;= 10<sup>5</sup></code></li> 
 <li>At most <code>2 * 10<sup>5</sup></code> calls will be made to <code>get</code> and <code>put</code>.</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>Hash Table | Linked List | Design | Doubly-Linked List</details><br>

<div>👍 20684, 👎 1022<span style='float: right;'><span style='color: gray;'><a href='https://github.com/labuladong/fucking-algorithm/discussions/939' target='_blank' style='color: lightgray;text-decoration: underline;'>bug 反馈</a> | <a href='https://labuladong.online/algo/fname.html?fname=jb插件简介' target='_blank' style='color: lightgray;text-decoration: underline;'>使用指南</a> | <a href='https://labuladong.online/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg' target='_blank' style='color: lightgray;text-decoration: underline;'>更多配套插件</a></span></span></div>

<div id="labuladong"><hr>

**通知：[新版网站会员](https://labuladong.online/algo/intro/site-vip/) 即将涨价！算法可视化编辑器上线，[点击体验](https://labuladong.online/algo/intro/visualize/)！**



<p><strong><a href="https://labuladong.online/algo/slug.html?slug=lru-cache" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

要让 `put` 和 `get` 方法的时间复杂度为 `O(1)`，我们可以总结出 `cache` 这个数据结构必要的条件：

1、显然 `cache` 中的元素必须有时序，以区分最近使用的和久未使用的数据，当容量满了之后要删除最久未使用的那个元素腾位置。

2、我们要在 `cache` 中快速找某个 `key` 是否已存在并得到对应的 `val`；

3、每次访问 `cache` 中的某个 `key`，需要将这个元素变为最近使用的，也就是说 `cache` 要支持在任意位置快速插入和删除元素。

哈希表查找快，但是数据无固定顺序；链表有顺序之分，插入删除快，但是查找慢，所以结合二者的长处，可以形成一种新的数据结构：哈希链表 `LinkedHashMap`：

![](https://labuladong.online/algo/images/LRU算法/4.jpg)

至于 `put` 和 `get` 的具体逻辑，可以画出这样一个流程图：

![](https://labuladong.online/algo/images/LRU算法/put.jpg)

根据上述逻辑写代码即可。

**详细题解：[算法就像搭乐高：带你手撸 LRU 算法](https://labuladong.online/algo/data-structure/lru-cache/)**

**标签：[数据结构](https://labuladong.online/algo/)，[设计](https://labuladong.online/algo/)**

## 解法代码

提示：🟢 标记的是我写的解法代码，🤖 标记的是 chatGPT 翻译的多语言解法代码。如有错误，可以 [点这里](https://github.com/labuladong/fucking-algorithm/issues/1113) 反馈和修正。

<div class="tab-panel"><div class="tab-nav">
<button data-tab-item="cpp" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">cpp🤖</button>

<button data-tab-item="python" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">python🤖</button>

<button data-tab-item="java" class="tab-nav-button btn active" data-tab-group="default" onclick="switchTab(this)">java🟢</button>

<button data-tab-item="go" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">go🤖</button>

<button data-tab-item="javascript" class="tab-nav-button btn " data-tab-group="default" onclick="switchTab(this)">javascript🤖</button>
</div><div class="tab-content">
<div data-tab-item="cpp" class="tab-item " data-tab-group="default"><div class="highlight">

```cpp
// 注意：cpp 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

#include <unordered_map>
#include <list>
#include <utility>

using namespace std;

class LRUCache {
private:
    int cap;
    unordered_map<int, pair<int, list<int>::iterator>> cache;
    list<int> l;

public:
    LRUCache(int capacity) {
        this->cap = capacity;
    }

    int get(int key) {
        if (cache.find(key) == cache.end()) {
            return -1;
        }
        // 将 key 变为最近使用
        makeRecently(key);
        return cache[key].first;
    }

    void put(int key, int val) {
        if (cache.find(key) != cache.end()) {
            // 修改 key 的值
            l.erase(cache[key].second);
        } else {
            if (l.size() >= this->cap) {
                // 链表头部就是最久未使用的 key
                cache.erase(l.front());
                l.pop_front();
            }
        }
        // 将新的 key 添加链表尾部
        l.push_back(key);
        cache[key] = {val, prev(l.end())};
    }

    void makeRecently(int key) {
        int val = cache[key].first;
        // 删除 key，重新插入到队尾
        l.erase(cache[key].second);
        l.push_back(key);
        cache[key] = {val, prev(l.end())};
    }
};
```

</div></div>

<div data-tab-item="python" class="tab-item " data-tab-group="default"><div class="highlight">

```python
# 注意：python 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
# 本代码已经通过力扣的测试用例，应该可直接成功提交。

class LRUCache:

    def __init__(self, capacity: int):
        self.cap = capacity
        self.cache = collections.OrderedDict()

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        # 将 key 变为最近使用
        self.makeRecently(key)
        return self.cache[key]

    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            # 修改 key 的值
            self.cache[key] = value
            # 将 key 变为最近使用
            self.makeRecently(key)
            return

        if len(self.cache) >= self.cap:
            # 链表头部就是最久未使用的 key
            oldestKey = next(iter(self.cache))
            self.cache.pop(oldestKey)
        # 将新的 key 添加链表尾部
        self.cache[key] = value

    def makeRecently(self, key: int) -> None:
        value = self.cache.pop(key)
        # 删除 key，重新插入到队尾
        self.cache[key] = value
```

</div></div>

<div data-tab-item="java" class="tab-item active" data-tab-group="default"><div class="highlight">

```java
class LRUCache {
    int cap;
    LinkedHashMap<Integer, Integer> cache = new LinkedHashMap<>();
    public LRUCache(int capacity) {
        this.cap = capacity;
    }

    public int get(int key) {
        if (!cache.containsKey(key)) {
            return -1;
        }
        // 将 key 变为最近使用
        makeRecently(key);
        return cache.get(key);
    }

    public void put(int key, int val) {
        if (cache.containsKey(key)) {
            // 修改 key 的值
            cache.put(key, val);
            // 将 key 变为最近使用
            makeRecently(key);
            return;
        }

        if (cache.size() >= this.cap) {
            // 链表头部就是最久未使用的 key
            int oldestKey = cache.keySet().iterator().next();
            cache.remove(oldestKey);
        }
        // 将新的 key 添加链表尾部
        cache.put(key, val);
    }

    private void makeRecently(int key) {
        int val = cache.get(key);
        // 删除 key，重新插入到队尾
        cache.remove(key);
        cache.put(key, val);
    }
}
```

</div></div>

<div data-tab-item="go" class="tab-item " data-tab-group="default"><div class="highlight">

```go
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

type LRUCache struct {
    capacity int
    cache    map[int]int
    keys     []int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]int),
        keys:     make([]int, 0),
    }
}

func (this *LRUCache) Get(key int) int {
    if val, ok := this.cache[key]; !ok {
        return -1
    } else {
        // 将 key 变为最近使用
        this.makeRecently(key)
        return val
    }
}

func (this *LRUCache) Put(key int, value int) {
    if _, ok := this.cache[key]; ok {
        // 修改 key 的值
        this.cache[key] = value
        // 将 key 变为最近使用
        this.makeRecently(key)
        return
    }

    if len(this.cache) >= this.capacity {
        // 链表头部就是最久未使用的 key
        oldestKey := this.keys[0]
        this.keys = this.keys[1:]
        delete(this.cache, oldestKey)
    }
    // 将新的 key 添加链表尾部
    this.cache[key] = value
    this.keys = append(this.keys, key)
}

func (this *LRUCache) makeRecently(key int) {
    val := this.cache[key]
    // 删除 key，重新插入到队尾
    delete(this.cache, key)
    this.cache[key] = val
    // Move the key to the end to mark it as recently used.
    this.keys = append(this.keys, key)
    // Remove the old occurrence of the key.
    index := 0
    for i, k := range this.keys {
        if k == key {
            index = i
            break
        }
    }
    this.keys = append(this.keys[:index], this.keys[index+1:]...)
}
```

</div></div>

<div data-tab-item="javascript" class="tab-item " data-tab-group="default"><div class="highlight">

```javascript
// 注意：javascript 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

var LRUCache = function(capacity) {
    this.cap = capacity;
    this.cache = new Map();
};

LRUCache.prototype.get = function(key) {
    if (!this.cache.has(key)) {
        return -1;
    }
    // 将 key 变为最近使用
    this.makeRecently(key);
    return this.cache.get(key);
};

LRUCache.prototype.put = function(key, val) {
    if (this.cache.has(key)) {
        // 修改 key 的值
        this.cache.set(key, val);
        // 将 key 变为最近使用
        this.makeRecently(key);
        return;
    }

    if (this.cache.size >= this.cap) {
        // 链表头部就是最久未使用的 key
        const oldestKey = this.cache.keys().next().value;
        this.cache.delete(oldestKey);
    }
    // 将新的 key 添加链表尾部
    this.cache.set(key, val);
};

LRUCache.prototype.makeRecently = function(key) {
    const val = this.cache.get(key);
    // 删除 key，重新插入到队尾
    this.cache.delete(key);
    this.cache.set(key, val);
};
```

</div></div>
</div></div>

**类似题目**：
  - [剑指 Offer II 031. 最近最少使用缓存 🟠](/problems/OrIXps)

</details>
</div>

