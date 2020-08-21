package leetcode

/****
给定一个无序的整数数组，找到其中最长上升子序列的长度。

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 */
func MaxSubSequece(nums []int) int {
	//dp[i] //以i结尾的最长子序列
	if(len(nums) <2){
		return len(nums)
	}
	dp := make([]int,len(nums)+1)
	dp[0] = 1;
	maxLen :=1;
	//maxVal := dp[0];
	for i:=1;i< len(nums);i++ {
		//maxVal 上一波的最大值
		if(nums[i-1] <= nums[i]){ //向上
			dp[i] = dp[i-1]+1;
			//maxLen++;
		}else{ //向下
			for j:=i-1;j>=0;j-- {
				if(nums[j]<nums[i]){
					dp[i] = dp[j]+1
					break
				}
				if(j==0){
					dp[i] = 1;
				}
			}
		}
		if maxLen < dp[i] {
			maxLen = dp[i]
		}

	}
	return maxLen;
}
