package piscine

func BTreeApplyPostorder(root *piscine.TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
	f(root.Data)

}