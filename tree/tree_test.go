package tree

import "testing"

var buildTreeCase = []struct {
	preOrderArr []int
	inOrderArr []int
}{
	{preOrderArr: []int{1,2}, inOrderArr: []int{2, 1}},
	{preOrderArr: []int{3,9,20,15,7}, inOrderArr: []int{9,3,15,20,7}},
}

func TestBuildTree(t *testing.T) {
	for _, ca := range buildTreeCase {
		BFSTree(buildTree(ca.preOrderArr, ca.inOrderArr))
		println()
	}
}

func preOrderTree(root *TreeNode) {
	if root == nil {
		print("null ")
		return
	}
	print(root.Val)
	print("\t")
	preOrderTree(root.Left)
	preOrderTree(root.Right)
}

func BFSTree(root *TreeNode) {
	BFSTreeSub([]*TreeNode{root})
}

func BFSTreeSub(root []*TreeNode) {
	subTreeNode := make([]*TreeNode, 0)
	nullCount := 0
	for _, r := range root {
		if r == nil {
			print("null	")
			nullCount ++
			continue
		}
		print(r.Val)
		print("\t")
		subTreeNode = append(subTreeNode, r.Left)
		subTreeNode = append(subTreeNode, r.Right)
	}
	if nullCount == len(root) {
		return
	}
	BFSTreeSub(subTreeNode)
}
