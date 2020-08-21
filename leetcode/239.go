package leetcode

import (
	"fmt"
	"math"
)

func MaxSlidingWindow(nums []int, k int) []int {
	i := 0;
	//m1 := make(map[int]int ,len(nums))
	maxV := []int{};
	s_maxV := 0;
	for j, _ :=  range nums {
		if j -i +1 < k {
			continue;
		}
		fmt.Println(nums[i:j+1],i,j)
		//if len(maxV) > 0 && maxV[len(maxV)-1] > nums[j] && nums[i] < maxV[len(maxV)-1]  {
		//	s_maxV =  maxV[len(maxV)-1];
		//}else {
		//	s_maxV = getmaxV(nums[i:j+1])
		//}
		s_maxV = getmaxV(nums[i:j+1])
		maxV = append(maxV,s_maxV);
		if j -i +1 == k {
			i++;
		}
	}
	return maxV;
}
func getmaxV(nums []int) int {
	MaxV:= math.MinInt16;
	for _,v := range nums {
		if MaxV < v {
			MaxV = v;
		}
	}
	return MaxV;
}
