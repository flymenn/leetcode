package leetcode
//爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。
//你有多少种不同的方法可以爬到楼顶呢？
//**注意：**给定 n 是一个正整数。
func ClimbStairs(n int) int {
	if(n <= 2) {
		return n;
	}
	dp := make([]int,n+1);
	dp[1] = 1;
	dp[2] = 2;
	for i:=3;i<=n;i++ {
		dp[i] = dp[i-1]+dp[i-2]
	}


	return dp[n];
}
