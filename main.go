package main

import (
	"algo/leetcode"
)

func main()  {
	//strs := []string{"flow","flower","fldd"}
	//str := leetcode.LongestCommonPrefix(strs)
	//fmt.Print(str);

	//prices:= []int{7,1,5,3,6,4}
	//prices:= []int{1,2,3,4,5}
	//ret:= leetcode.MaxProfit(prices)
	//fmt.Println("-------");
	//fmt.Println(ret);


	//ret := leetcode.ClimbStairs(5);
	//fmt.Println(ret);


	//nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	//ret := leetcode.MaxSum(nums);
	//fmt.Println(ret)


	//nums := []int{10,9,2,5,3,7,101,18}
	//ret := leetcode.MaxSubSequece(nums);
	//fmt.Println(ret)

	//nums := [][]int{
	//	{2},
	//	{3,4},
	//	{6,5,7},
	//	{4,1,8,3},
	//}
	//ret := leetcode.MinimumTotal(nums);
	//fmt.Println(ret)


	//a := [3]int{0,1,2}
	//s := a[1:2]
	//s[0] =11
	//s = append(s,12);
	//s = append(s,13);
	//s[0] =21
	//a[1] = 30;
	//fmt.Println(a)
	//fmt.Println(s)
	//



	node := &leetcode.TreeNode{Value: 1}
	node.Left = &leetcode.TreeNode{Value: 2}
	node.Rigth = &leetcode.TreeNode{Value: 3}
	node.Left.Left = &leetcode.TreeNode{Value: 4}
	node.Left.Rigth = &leetcode.TreeNode{Value: 5}
	node.Rigth.Left = &leetcode.TreeNode{Value: 6}
	node.Rigth.Rigth = &leetcode.TreeNode{Value: 7}
	//ret := leetcode.Dfs(node)

	//ret := leetcode.MaxDepth(node)

	//ret := leetcode.LevelTree(node)

	//fmt.Println(ret);
	//str := "abcabcbb";
	////ret := leetcode.LongSubStr(str)
	//ret := leetcode.LengthOfLongestSubstring(str)
	//fmt.Println(ret)	\
	//
	//
	//
	//nums := []int {1,-1};
	//
	//ret := leetcode.MaxSlidingWindow(nums,1)
	//fmt.Println(ret)
	//nums := []int {1,8,6,2,5,4,8,3,7};
	//
	//ret := leetcode.MaxArea(nums)
	//fmt.Println(ret)

	//nums := []int {-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6};
	//
	//ret := leetcode.ThreeSum(nums)
	//fmt.Println(ret)






}