package piscine

func BTreeApplyByLevel(root *TreeNode, fn interface{})  {
if root == nil {
		return
	}
	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}