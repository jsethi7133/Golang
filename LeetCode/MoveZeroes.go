func moveZeroes(nums []int)  {
    lastNonZeroPos:=0
    for i:=0;i<len(nums);i++{
        if nums[i]!=0{
            nums[lastNonZeroPos]=nums[i]
            lastNonZeroPos+=1
        }
    }
    for i:=lastNonZeroPos;i<len(nums);i++{
        nums[i]=0
    }
}
