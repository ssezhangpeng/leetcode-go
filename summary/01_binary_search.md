## 二分搜索专题
### 题目汇总
- leetcode_0033
- leetcode_0034
- leetcode_0035
- leetcode_0069
- leetcode_0081
- leetcode_0053
- leetcode_0154
- leetcode_0374
- leetcode_0540
- leetcode_0704
- leetcode_0852
- leetcode_0875(UNDO)

### 解题模版
```go
// leetcode_0035
func searchInsert(nums []int, target int) int {
    size := len(nums);

    // 此处选择搜索区间为左闭右闭区间
    l, r := 0, size - 1
    // 当 l == r 时，[l, r]是一个有效区间
    for l <= r {
        mid := l + (r - l) >> 1
        if nums[mid] == target {
            return mid
        } else if (nums[mid] < target) {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return l
}
```

### 专题总结
[Leetcode总结篇](https://leetcode-cn.com/problems/search-insert-position/solution/te-bie-hao-yong-de-er-fen-cha-fa-fa-mo-ban-python-/)