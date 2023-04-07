func twoSum(nums []int, target int) []int {
    //return bruteForce(nums, target)
    return onePassHash(nums, target)
}

func bruteForce(nums []int, target int) []int{
    var res []int
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]+nums[j]==target{
                res=append(res,i,j)
                return res
            }

        }
    }
    return nil
}

func onePassHash(nums []int, target int) []int{
    resMap:=make(map[int]int)
    var res []int
    for i:=0;i<len(nums);i++{
        if val,ok:=resMap[target-nums[i]];ok{
            res=append(res,i,val) 
            return res
        }
        resMap[nums[i]]=i
    }
    return nil
}