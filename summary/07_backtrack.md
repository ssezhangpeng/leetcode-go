## 回溯专题
### 题目汇总
- leetcode_0017
- leetcode_0022
- leetcode_0037
- leetcode_0039
- leetcode_0040
- leetcode_0046
- leetcode_0047
- leetcode_0051
- leetcode_0052
- leetcode_0077
- leetcode_0078
- leetcode_0090
- leetcode_0093
- leetcode_0112
- leetcode_0131
- leetcode_0216
- leetcode_0257
- leetcode_0301(THINKING)
- leetcode_0491
- leetcode_0494

### 使用场景
- 组合问题：N个数里面按一定规则找出k个数的集合
- 切割问题：一个字符串按一定规则有几种切割方式
- 子集问题：一个N个数的集合里有多少符合条件的子集
- 排列问题：N个数按一定规则全排列，有几种排列方式
- 棋盘问题：N皇后，解数独等等


### 解题模版
```go
// leetcode_0040
func combinationSum2(candidates []int, target int) [][]int {
    var path []int
    var ans [][]int

    sort.Ints(candidates)

    backtracking(candidates, 0, 0, target, path, &ans)

    return ans
}

func backtracking(candidates []int, start, sum, target int, path []int, ans *[][]int) {
    // 1. 递归出口
    if sum >= target {
        if sum == target {
            temp := make([]int, len(path))
            copy(temp, path)
            *ans = append(*ans, temp)
        }
        return
    }

    // 2. 横向遍历每一个起点
    for i:=start; i<len(candidates); i++ {
        if i>start && candidates[i] == candidates[i-1] {
            continue
        }
        // 3. 尝试加入新元素
        path = append(path, candidates[i])
        sum += candidates[i]
        // 4. 纵向遍历到下一层(确定下一层的起点)
        backtracking(candidates, i+1, sum, target, path, ans)
        // 5. 返回上层，恢复现场
        sum -= candidates[i]
        path = path[:len(path)-1]
    }
}
```

### 专题总结
- [代码随想录](https://leetcode-cn.com/problems/permutations/solution/dai-ma-sui-xiang-lu-dai-ni-xue-tou-hui-s-mfrp/)