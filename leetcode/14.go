package leetcode

import "strings"

func LongestCommonPrefix(strs []string)  string {
	ret := ""
	if len(strs) == 0{
		return ret
	}
	ret = strs[0];
	for _,v := range strs{
		for len(ret)>0  {
			if(strings.Index(v,ret) == 0 ){ //是前缀
				break;
			}else {
				ret = ret[:len(ret)-1]
			}
		}

	}
	return ret
}