package leetcode

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0;
	}
	return max(MaxDepth(root.Left),MaxDepth(root.Rigth))+1
}

func  max(a int , b int ) int  {
	if a > b {
		return a;
	}
	return b;
}