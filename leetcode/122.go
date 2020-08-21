package leetcode

import "fmt"

func MaxProfit(prices []int) int {
	profit := 0;
	if(len(prices) == 0){
		return profit;
	}
	tence := 0; //1 up -1 down
	hasBuy:= 0;
	buyPoint := 0
	for k,v:= range prices{
		//fmt.Print(k);
		if k== 0  {
			continue;
		}

		if( prices[k-1] <= v )  { //网上走
			if(hasBuy == 0){ //没买过
				tence=1;
				buyPoint = k-1;
				hasBuy=1
			}
		} else if prices[k-1] >= v  && tence!= 0 { //往下走 && 原来的趋势不想下
			if(hasBuy == 1) { //买过
				tence=-1;
				profit = profit +  prices[k-1] - prices[buyPoint];
				hasBuy = 0;
			}
		}
		if(hasBuy == 1 && k == len(prices)-1){
			profit = profit +  prices[k] - prices[buyPoint];
			fmt.Print(profit);
		}

	}
	return profit
}
