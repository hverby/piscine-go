package piscine 

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	f(root.Left.Data)
	f(root.Data)
	f(root.Right.Left.Data)
	f(root.Right.Data)
}