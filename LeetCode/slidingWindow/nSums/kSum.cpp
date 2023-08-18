vector<vector<int>> nSumTarget(
    vector<int>& nums, int n, int start, int target) {

    int sz = nums.size();
    vector<vector<int>> res;
    // 至少是 2Sum，且数组大小不应该小于 n
    if (n < 2 || sz < n) return res;
    // 2Sum 是 base case
    if (n == 2) {
        // 双指针那一套操作
        int lo = start, hi = sz - 1;
        while (lo < hi) {
            int sum = nums[lo] + nums[hi];
            int left = nums[lo], right = nums[hi];
            if (sum < target) {
                while (lo < hi && nums[lo] == left) lo++;
            } else if (sum > target) {
                while (lo < hi && nums[hi] == right) hi--;
            } else {
                res.push_back({left, right});
                while (lo < hi && nums[lo] == left) lo++;
                while (lo < hi && nums[hi] == right) hi--;
            }
        }
    } else {
        // n > 2 时，递归计算 (n-1)Sum 的结果
        for (int i = start; i < sz; i++) {
            vector<vector<int>> 
                sub = nSumTarget(nums, n - 1, i + 1, target - nums[i]);
            for (vector<int>& arr : sub) {
                // (n-1)Sum 加上 nums[i] 就是 nSum
                arr.push_back(nums[i]);
                res.push_back(arr);
            }
            while (i < sz - 1 && nums[i] == nums[i + 1]) i++;
        }
    }
    return res;
}

// 实际上就是把之前的题目解法合并起来了，n == 2 时是 twoSum 的双指针解法，n > 2 时就是穷举第一个数字，然后递归调用计算 (n-1)Sum，组装答案。

// 根据之前几道题的时间复杂度可以推算，本函数的时间复杂度应该是 O(N^(n-1))，N 为数组的长度，n 为组成和的数字的个数。
// 需要注意的是，调用这个 nSumTarget 函数之前一定要先给 nums 数组排序


// https://appktavsiei5995.pc.xiaoe-tech.com/detail/i_62987931e4b0cedf38ba0684/1?from=p_629871eee4b01a48520729f7&type=6&parent_pro_id=