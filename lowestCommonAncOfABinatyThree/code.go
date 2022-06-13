package lowestCommonAncOfABinatyThree

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ancestor *TreeNode
	if root == nil{
		return nil
	}
	ancestor = lowestCommonAncestor(root.Left, p, q)
	if ancestor != nil{
		return ancestor
	}
	ancestor = lowestCommonAncestor(root.Right, p, q)
	if ancestor != nil{
		return ancestor
	}
	if find(root,p.Val) && find(root,q.Val){
		return root
	}
	return nil

}

func find(root *TreeNode, element int ) bool{
	if root == nil{
		return false
	}
	if root.Val == element{
		return true
	}
	return find(root.Left, element) || find(root.Right, element)
}
