package leetcode_0113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var path []int
	var ans [][]int
	pathSumCore(root, targetSum, path, &ans)
	return ans
}

func pathSumCore(root *TreeNode, targetSum int, path []int, ans *[][]int) {
	if root == nil {
		return
	}

	targetSum -= root.Val
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}

	pathSumCore(root.Left, targetSum, path, ans)
	pathSumCore(root.Right, targetSum, path, ans)

	path = path[:len(path)-1]
}
