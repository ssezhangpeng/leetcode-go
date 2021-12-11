## 深度优先搜索专题
### 题目汇总

- leetcode_0130
- leetcode_0200
- leetcode_0329(TIMEOUT)
- leetcode_0417
- leetcode_0463
- leetcode_0695
- leetcode_0980


### 解题模版
```go
// leetcode_0130
func solve(board [][]byte)  {
    m, n := len(board), len(board[0])

    // 遍历每个边缘的 ‘O’, 并以此作为起点进行 DFS
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if (i == 0 || i == m-1 || j == 0 || j == n - 1) && board[i][j] == 'O' {
                dfs(&board, i, j)
            }
        }
    }

    // 重新遍历每个位置,把没有在 DFS 中遍历到的都标记为 ‘X’
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if board[i][j] == '#' {
                board[i][j] = 'O'
            } else {
                board[i][j] = 'X'
            }
        }
    }

    return
}

func dfs(board *[][]byte, i, j int) {
    m, n := len(*board), len((*board)[0])
    // 递归出口
    if i < 0 || i >= m || j < 0 || j >= n || (*board)[i][j] != 'O' {
        return
    }

    (*board)[i][j] = '#'
    dfs(board, i-1, j)
    dfs(board, i+1, j)
    dfs(board, i, j-1)
    dfs(board, i, j+1)
}
```

### 专题总结