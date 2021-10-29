package leetcode0337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var f = make(map[*TreeNode]int)
var g = make(map[*TreeNode]int)

func rob(root *TreeNode) int {
	dfs(root)
	return max(f[root], g[root])
}

func dfs(root *TreeNode) {
	// 递归出口
	if root == nil {
		return
	}

	dfs(root.Left)
	dfs(root.Right)

	// 选择 root 节点
	f[root] = root.Val + g[root.Left] + g[root.Right]

	// 不选择 root 节点
	g[root] = max(f[root.Left], g[root.Left]) + max(f[root.Right], g[root.Right])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
