package leetcode

import (
	"fmt"
)

func MaxArea(height []int) int {

	i,j := 0,len(height)-1;
	MaxSqure := 0;
	for i<j {

		mi := height[i] - height[j];
		fmt.Println(MaxSqure,i,j,mi)
		if mi > 0 {
			MaxSqure  = max(height[j] * (j-i),MaxSqure)
			j--;
		}else{
			MaxSqure  = max(height[i] * (j-i),MaxSqure)
			i++;
		}
	}
	return MaxSqure;
}

