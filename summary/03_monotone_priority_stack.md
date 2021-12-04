## 单调栈专题
### 题目汇总
- leetcode_0042
- leetcode_0084
- leetcode_0239
- leetcode_0496
- leetcode_0402
- leetcode_0316
- leetcode_0321
- leetcode_0503
- leetcode_0739

### 解题模版
```go
// leetcode_0042
func trap(height []int) int {
	var water int
	var stack []int

	for i:=0; i<len(height); i++ {
		// 尝试持续出栈
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				continue
			}

			left := stack[len(stack)-1]
			h := min(height[i], height[left]) - height[top]
			w := i - left - 1

			water += h * w
		}
		stack = append(stack, i)
	}

	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### 专题总结
- [wg讲算法](https://www.bilibili.com/video/BV1MA411T7YR?from=search&seid=12857983768183582648&spm_id_from=333.337.0.0)
- [古城算法](https://www.bilibili.com/video/BV17z4y1y7tS?spm_id_from=333.999.0.0)