## 常用算法专题
### 求最小公约数(gcd)
```go
func GDC(a, b int) int {
	if a%b == 0 {
		return b
	}
	return GDC(b, a%b)
}
```
### 求最小公倍数(lcm)
```go
func LCM(a, b int) int {
	return a * b / GDC(a, b)
}
```