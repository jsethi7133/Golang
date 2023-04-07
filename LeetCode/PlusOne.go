func plusOne(digits []int) []int {
    for i:=len(digits)-1;i>=0;i--{
        if digits[i]==9{
            digits[i]=0
        }else{
            digits[i]=digits[i]+1
            break
        }
    }
    if digits[0]==0{
        digitsNew:=[]int{1}
        digitsNew=append(digitsNew,digits...)
        return digitsNew
    }
    return digits
}
