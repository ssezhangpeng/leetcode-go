## 常用算法专题
### 求最小公约数(gcd)
```go
func GDC(a, b int) int {
	if a % b == 0 {
		return b
	}
	return GDC(b, a % b)
}
```
### 求最小公倍数(lcm)
```go
func LCM(a, b int) int {
	return a * b / GDC(a, b)
}
```

### TopK 问题
```go
func TopK(nums []int, k int) []int {
	return nil
}
```

### KMP 问题
```go
- leetcode_0028
- leetcode_0456
https://www.programmercarl.com/0028.实现strStr.html#其他语言版本
```

### 洗牌算法 问题
```go
- leetcode_0384
https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
```