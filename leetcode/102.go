package leetcode

//给定一个二叉树，返回其按层次遍历的节点值。（即逐层地，从左到右访问所有节点）。
/**
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：[[3],[9,20],[15,7]]
*/

func LevelTree(root *TreeNode) [][]int {
	return dfs(root,0,[][]int{});
}

func dfs(root *TreeNode,level int ,ret [][]int) [][]int {
	if(root == nil){
		return ret;
	}
	if(level == len(ret)){
		ret = append(ret,[]int{});
	}
	ret[level] = append(ret[level],root.Value)
	ret = dfs(root.Left,level+1,ret);
	ret = dfs(root.Rigth,level+1,ret);
	return ret
}

//func levelTree2(root *TreeNode) [][]int  {
//	//queue := list.New()
//}
