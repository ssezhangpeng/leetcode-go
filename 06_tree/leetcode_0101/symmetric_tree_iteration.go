package leetcode_0101

func isSymmetricIteration(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricIterationCore(root, root)
}

func isSymmetricIterationCore(p *TreeNode, q *TreeNode) bool {
	queue := make([]*TreeNode, 0)

	queue = append(queue, p)
	queue = append(queue, q)

	for len(queue) > 0 {
		p = queue[0]
		queue = queue[1:]

		q = queue[0]
		queue = queue[1:]

		if p == nil && q == nil {
			continue
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		queue = append(queue, p.Left)
		queue = append(queue, q.Right)

		queue = append(queue, p.Right)
		queue = append(queue, q.Left)
	}

	return true
}
