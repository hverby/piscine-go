package piscine

import "fmt"

func BTreeApplyByLevel(root *TreeNode, fn interface{})  {
if root == nil {
		return
	}
	fmt.Println(fn.Data)
	BTreeApplyPreorder(root.Left, fn)
	BTreeApplyPreorder(root.Right, fn)
}