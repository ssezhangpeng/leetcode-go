package leetcode_0133

type Node struct {
	Val       int
	Neighbors []*Node
}

var (
	// key: node, value: cloneNode
	m = make(map[*Node]*Node)
)

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// 如果该节点已经访问过，则直接从 map 中取回对应的克隆节点返回
	if v, ok := m[node]; ok {
		return v
	}

	// 克隆节点，但不会克隆该节点的邻居节点
	cloneNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}

	// 存入哈希表
	m[node] = cloneNode

	// 遍历该节点的所有邻居节点，并更新克隆节点的邻居节点
	for i := 0; i < len(node.Neighbors); i++ {
		cloneNode.Neighbors[i] = cloneGraph(node.Neighbors[i])
	}

	return cloneNode
}
