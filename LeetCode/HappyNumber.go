func isHappy(n int) bool {
    hMap:= make(map[int]int)
    for (n!=1 && hMap[n]!=1){
        hMap[n]=1
        n=getNext(n)
    }
    return n==1
}

func getNext(n int) int{
    totalSum:=0
    for(n>0){
        d:=n%10
        n=n/10
        totalSum=totalSum+d*d
    }
    return totalSum
}
