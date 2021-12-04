## 双指针专题
### 题目汇总
- leetcode_0019
- leetcode_0141
- leetcode_0142
- leetcode_0202
- leetcode_0234
- leetcode_0287
- leetcode_0876

### 解题模版
```go
// leetcode_0876
func middleNode(head *ListNode) *ListNode {
    // 可以通过调节 fast 的初始位置, 确定中间指针的位置是偏前或者偏后
    slow, fast := head, head

    // 固定条件循环条件
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}
```

### 专题总结
[Leetcode总结篇](https://leetcode-cn.com/problems/find-the-duplicate-number/solution/qian-duan-ling-hun-hua-shi-tu-jie-kuai-man-zhi-z-3/)