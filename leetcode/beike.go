package leetcode

func getO(n int)  {


	visit :=  make([]int,n)
	_getO(n,visit);
}


func _getO(n int ,visit []int)[]int {
	if n%2 == 0 || n%3 == 0 || n%5 == 0 || n%7 == 0 {
		return visit;
	}
	i:= 2;
	value := 2;
	for(true ){
		if(value<n){
			value = i<<2;
			visit = append(visit,value);
		}
		i++;
	}
	//dp[i] //以他结尾的素数
	return visit;
}


