package leetcode

import (
	"fmt"
	"math"
)

/**
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]

 */
func MinimumTotal(triangle [][]int) int {
	y :=  len(triangle)
	if y==1 {
		return triangle[0][0];
	}
	if(y == 0) {
		 return 0
	}
	dp:= make([][]int,y);
	for k,v := range triangle{
		dp[k] = make([]int,len(v))
	}
	dp[0][0] = triangle[0][0];
	min := 0;
	for k,v := range triangle{
		if(k== 0) {
			continue;
		}
		for kk,vv := range v{
			if(kk == 0){
				min = dp[k-1][kk]
			}else if( kk < len(v)-1){
				min = minValue(dp[k-1][kk],dp[k-1][kk-1])
			}else if kk == (len(v)-1) {
				min = dp[k-1][kk-1];
			}else{
				break;
			}
			dp[k][kk] =  min  + vv;
		}
	}
	fmt.Println(dp);
	ret  := math.MaxInt64;
	for _,v := range dp[y-1] {
		ret = minValue(v,ret);
	}
	return  ret;
}

func minValue(a int, b int )  int {
	if a< b {
		return a;
	}else{
		return b;
	}
}
