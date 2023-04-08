func majorityElement(nums []int) int {
    majorityCount:=int(len(nums)/2)
    for i:=0;i<len(nums);i++{
        count:=0
        for j:=0;j<len(nums);j++{
            if(nums[i]==nums[j]){
                count=count+1
            }
        }
        if(count>majorityCount){
            return nums[i]
        }
    }
    return -1
}
