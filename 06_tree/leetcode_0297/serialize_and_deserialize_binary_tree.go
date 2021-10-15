package leetcode_0297

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 以 '#' 标识空节点，以 ',' 表示一个节点值的结束
	ans := ""
	this.serializeCore(root, &ans)
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	u := 0
	return this.deserializeCore(data, &u)
}

func (this *Codec) serializeCore(root *TreeNode, ans *string) {
	if root == nil {
		*ans = *ans + "#" + ","
		return
	}

	*ans = *ans + strconv.Itoa(root.Val) + ","

	this.serializeCore(root.Left, ans)
	this.serializeCore(root.Right, ans)
}

func (this *Codec) deserializeCore(s string, i *int) *TreeNode {
	if s[*i] == '#' {
		*i = *i + 2
		return nil
	}

	value, flag := 0, 1
	if s[*i] == '-' {
		flag = -1
		*i = *i + 1
	}
	for s[*i] != ',' {
		value = value*10 + int(s[*i]-'0')
		*i = *i + 1
	}
	// 跳过 ','
	*i = *i + 1
	value = value * flag

	root := &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}

	root.Left = this.deserializeCore(s, i)
	root.Right = this.deserializeCore(s, i)

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
