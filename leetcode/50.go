package leetcode

func MaxSum(nums []int) int {
	//dp[i] 表示以i结尾的最大和
	dp := make([]int,len(nums)+1);
	dp[0] = nums[0];
	maxVal := dp[0];
	for i:=1;i<len(nums);i++ {
		if(dp[i-1]<0){
			dp[i] = nums[i]
		}else{
			dp[i] = nums[i]+dp[i-1]
		}
		//fmt.Println(dp[i])
		if(dp[i]>maxVal){
			maxVal = dp[i]
		}
	}
	return maxVal
}