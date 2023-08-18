```js
int left = 0, right = 0;

while (right < s.size()) {
    // 增大窗口
    window.add(s[right]);
    right++;
    
    while (window needs shrink) {
        // 缩小窗口
        window.remove(s[left]);
        left++;
    }
}

```

- 虽然滑动窗口代码框架中有一个嵌套的 while 循环，但算法的时间复杂度依然是 O(N)，其中 N 是输入字符串/数组的长度。
    - 为什么呢？简单说，字符串/数组中的每个元素都只会进入窗口一次，然后被移出窗口一次，不会说有某些元素多次进入和离开窗口，所以算法的时间复杂度就和字符串/数组的长度成正比