package leetcode

//深度优先算法


type TreeNode struct {
	Value int
	Left,Rigth *TreeNode
}

func Dfs(root *TreeNode) []int{
	var stack []*TreeNode;
	var res []int;
	//var cur TreeNode;
	stack = append(stack,root);
	for (len(stack)>0){
		cur := stack[len(stack)-1];
		stack = stack[:len(stack)-1]
		if cur.Rigth !=nil{
			stack = append(stack,cur.Rigth)
		}
		if cur.Left !=nil {
			stack = append(stack,cur.Left)
		}

		res = append(res,cur.Value)
	}
	return res;
}
