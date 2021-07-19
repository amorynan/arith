package tree
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildSubTree(preorder, inorder, 0, 0, len(preorder)-1)
}

func buildSubTree(preorder, inorder []int, pi, is, ie int) *TreeNode {
	if ie == is {
		return &TreeNode{
			Val: inorder[is],
		}
	}

	head := &TreeNode{
		Val : preorder[pi],
	}
	for i := is; i <= ie; i ++{
		if inorder[i] == preorder[pi] {
			// find root , and split left & right tree
			// left tree
			if i != is {
				head.Left = buildSubTree(preorder, inorder, pi+1, is, i-1)
			}
			// right tree
			if i != ie {
				head.Right = buildSubTree(preorder, inorder, pi+(i-is)+1, i+1, ie)
			}
		}
	}
	return head
}
