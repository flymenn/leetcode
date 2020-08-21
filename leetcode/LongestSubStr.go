package leetcode

import "fmt"

func LongSubStr(str string) string {
	//w := []string{};
	//w =append(w,string(str[0]))
	left,right :=0,0; //坐标
	Index := make(map [string]int,len(str))
	Index[string(str[0])] = 0;
	Maxleft,MaxRight := 0,0;
	for i:=1;i<len(str);i++ {
		fmt.Println(left,right,Maxleft,MaxRight)
		right = i;
		if befor_have, ok := Index[string(str[i])];ok {
			if(befor_have > left && right-left-1 > MaxRight-Maxleft){
				Maxleft = left;
				MaxRight = right-1;
				left = befor_have+1;
			}
			fmt.Println(befor_have,string(str[i]))

		}
		Index[string(str[i])] = i;
	}
	if right-left > MaxRight-Maxleft{
		Maxleft = left;
		MaxRight = right;
	}
	return str[Maxleft:MaxRight+1]
}

func LengthOfLongestSubstring2(s string) int {
	if len(s)<=1 {
		return len(s)
	}
	left,right :=0,0; //坐标
	Index := make(map [string]int,len(s))
	Index[string(s[0])] = 0;
	MaxLengh := 0
	for i:=1;i<len(s);i++ {
		fmt.Println(left,right,MaxLengh,Index)
		if befor_have, ok := Index[string(s[i])];ok {
			if befor_have >= left  {
				MaxLengh = getMax(right-left+1,MaxLengh);
				//fmt.Println(MaxLengh);
				left = befor_have+1;
			}
		}
		right = i;
		Index[string(s[i])] = i;
	}
	MaxLengh = getMax(right-left+1,MaxLengh);
	return MaxLengh;
}
func getMax(a, b int ) int {
	if a> b {
		return a
	}
	return b
}