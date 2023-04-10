func gcdOfStrings(str1 string, str2 string) string {
    l1:=len(str1)
    l2:=len(str2)
    for i:=int(math.Min(float64(l1),float64(l2)));i>0;i--{
        if valid(str1,str2,i,l1,l2){
            return str1[:i]
        }
    }
    return ""
}

func valid(str1 string, str2 string, i int, l1 int, l2 int) bool{
    if(l1%i>0 || l2%i>0){
        return false
    }else{
        base:=str1[:i]
        return len(strings.Replace(str1, base, "", -1))==0 && len(strings.Replace(str2, base, "", -1))==0
    }
}
