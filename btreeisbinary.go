package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if BTreeIsBinary(root.Left) && root.Data > root.Left.Data {
		return false
	}
	if BTreeIsBinary(root.Right) && root.Data < root.Right.Data {
		return false
	}

	if !BTreeIsBinary(root.Left) || BTreeIsBinary(root.Right){
		return false
	}
	return true
}