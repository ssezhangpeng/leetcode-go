## 滑动窗口专题
### 题目汇总
- leetcode_0003
- leetcode_0076
- leetcode_0159(PLUS)
- leetcode_0209
- leetcode_0340(UNDO)
- leetcode_0395(UNDO)
- leetcode_0424(UNDO)
- leetcode_0992(UNDO)
- leetcode_1248(UNDO)

### 解题模版
```go
// leetcode_0076
func minSubArrayLen(target int, nums []int) int {
	size := len(nums)
	var sum int

	minLength := 1 << 15
	left, right := 0, 0
	// 1. 持续右移 right
	for right=0; right < size; right ++ {
		sum += nums[right]
		// 2. 持续左移 left
		for sum >= target {
			currLength := right - left + 1
			// 3. 处理满足条件的结果
			if currLength < minLength {
				minLength = currLength
			}
			// 4. 找到部分结果后，继续向后寻找
			sum -= nums[left]
			left++
		}
	}

	if minLength == 1 << 15 {
		return 0
	}

	return minLength
}
```

### 专题总结
- [古城算法](https://www.bilibili.com/video/BV1PU4y147tP?from=search&seid=381302338992998897&spm_id_from=333.337.0.0)