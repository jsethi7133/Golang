func searchRange(nums []int, target int) []int {
    var res []int
    for i:=0;i<len(nums);i++{
        if nums[i]==target{
            res=append(res,i)
            break
        }
    }
    for i:=len(nums)-1;i>=0;i--{
        if nums[i]==target{
            res=append(res,i)
            break
        }
    }
    if len(res)!=0{
    return res
    }else{
        return append(res,-1,-1)
    }
}
