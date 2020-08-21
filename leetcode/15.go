package leetcode

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums);
	ret := [][]int{};
	sum := 0;

	for k,v := range  nums{
		j := len(nums)-1;
		two_sum := sum- v;

		for i:= k+1;i<j; {
			fmt.Println(two_sum,k,i,j);
			if (nums[i]+ nums[j]) >two_sum {
				j--;
			}else if (nums[i]+ nums[j]) < two_sum  {
				i++;
			}else{ //相等
				if !isContain(ret,[]int{v,nums[i],nums[j]}) {
					ret = append(ret,[]int{v,nums[i],nums[j]})
				}
				j--;
				i++;
				//break;
			}
		}
	}
	return ret;
}

func isContain(num1 [][]int,num2 []int)  bool {
	for _,v := range num1{
		for kk,vv :=  range v {
			if vv != num2[kk] {
				goto END;
			}
		}
		return true;
		END:
	}
	return false;
}
