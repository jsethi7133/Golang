func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1=append(nums1[:m],nums2[:n]...)
    sort.Ints(nums1)
    //bubbleSort(nums1)
}

func bubbleSort(nums []int){
    for i:=0;i<len(nums)-1;i++{
        for j:=0;j<len(nums)-i-1;j++{
            if(nums[j]>nums[j+1]){
                nums[j], nums[j+1]=nums[j+1], nums[j]
            }
        }
    }
}
