package leetcode

import (
	"math"
)

func LengthOfLongestSubstring(s string) int {
	if len(s) <=1 {
		return len(s)
	}
	m1 := make(map [int32]int,len(s))
	i := 0;
	maxLen := 0;
	for k,v := range s{
		if k  == 0 {
			m1[v]  = 0;
			continue;
		}
		if hava_k,ok := m1[v];ok {
			if hava_k >= i {
				maxLen = int(math.Max(float64(maxLen),float64(k-i)));
				i = hava_k +1;
			}
		}
		m1[v] = k;
		if k == len(s) -1 {
			maxLen = int(math.Max(float64(maxLen),float64(k - i +1 )));
		}
	}
	return maxLen
}