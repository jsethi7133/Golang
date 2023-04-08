func moveZeroes(nums []int)  {
    var res []int
    for i:=0;i<len(nums);i++{
        if(nums[i]!=0){
            res=append(res,nums[i])
        }
    }
    fmt.Println(res)
    for j:=0;j<len(nums)-len(res)+1;j++{
        res=append(res,0)
    }
    fmt.Println(res)
    for k:=0;k<len(nums);k++{
        nums[k]=res[k]
    }
}
