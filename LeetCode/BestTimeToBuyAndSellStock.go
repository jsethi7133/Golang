func maxProfit(prices []int) int {
    //return bruteF(prices)
    return singleP(prices)
}

func bruteF(prices []int) int{
        maxProfit:=0
    for i:=0;i<len(prices)-1;i++{
        for j:=i+1;j<len(prices);j++{
            profit:=prices[j]-prices[i]
            if profit>maxProfit{
                maxProfit=profit
            }
        }
    }
    return maxProfit
}

func singleP(prices []int) int{
    minPrice:=math.MaxInt
    maxProfit:=0
    for i:=0;i<len(prices);i++{
        if(prices[i]<minPrice){
            minPrice=prices[i]
        }else if(prices[i]-minPrice>maxProfit){
            maxProfit=prices[i]-minPrice
        }
    }
    return maxProfit
}
