func containsDuplicate(nums []int) bool {
    //return bruteF(nums)
    return hashM(nums)
}

func bruteF(nums []int) bool{
        count:=0
    for i:=0;i<len(nums);i++{
        count=0
        for j:=0;j<len(nums);j++{
            if nums[i]==nums[j]{
                count++
                if(count>1){
                    return true
                }
            }
        }
    }
    return false
}

func hashM(nums []int) bool{
        res:=make(map[int]int, len(nums))
    for _,n:=range nums{
        if _,ok:=res[n];ok{
            return true
        }
        res[n]=1
    }
    return false
}
