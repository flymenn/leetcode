package leetcode

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums);
	minv := math.MaxInt64;
	ret := 0;
	for k,v := range nums{
		j := len(nums)-1;
		two_sum := target-v;

		for i:=k+1; i<j ; {
			diff  := (nums[i] +nums[j])-two_sum
			if  diff > 0  {
				minv = minAbs(minv,diff);
				j--;
			}else if diff < 0{
				minv = minAbs(minv,diff);
				i++;
			}else{
				return target;
			}
			if minv == diff {
				ret = diff+target;
			}
		}
	}
	return ret;

}

func minAbs(a int, b int )  int   {
	if math.Abs(float64(a)) < math.Abs(float64(b)){
		return a;
	}
	return b;
}