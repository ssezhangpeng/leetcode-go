## 广度优先搜索专题
### 题目汇总
- leetcode_0102
- leetcode_0111
- leetcode_0127
- leetcode_0130
- leetcode_0207
- leetcode_0490(PLUS)
- leetcode_0505(PLUS)
- leetcode_0542 [Trick](https://leetcode-cn.com/problems/01-matrix/solution/01ju-zhen-by-leetcode-solution/)
- leetcode_0994 [Trick](https://leetcode-cn.com/problems/rotting-oranges/solution/fu-lan-de-ju-zi-by-leetcode-solution/)

### Tricks
- [多源BFS](https://leetcode-cn.com/problems/rotting-oranges/solution/fu-lan-de-ju-zi-by-leetcode-solution/)


### 解题模版
```go
// leetcode_0111
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var depth int
    var current []*TreeNode
    // 1. 初始化第一层
    current = append(current, root)
    // 2. 对每一圈波纹进行遍历(由内到外)
    for len(current) > 0 {
        depth++
        // 3. 暂存下一圈波纹的信息
        var next []*TreeNode
        // 4. 对某一圈波纹进行遍历(旋转一圈,收集下一圈的信息)
        for len(current) > 0 {
            node := current[0]
            current = current[1:]
            
            if node.Left == nil && node.Right == nil {
                return depth
            }

            if node.Left != nil {
                next = append(next, node.Left)
            }
            if node.Right != nil {
                next = append(next, node.Right)
            }
        }
        // 5. 将下一圈的信息赋值给新的圈,继续下一圈的处理
        current = next
    }

    return depth
}
```

### 专题总结
- [古城算法](https://www.bilibili.com/video/BV1Rz4y1Z7tJ?spm_id_from=333.999.0.0)