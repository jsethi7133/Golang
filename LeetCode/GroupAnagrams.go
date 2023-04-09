func groupAnagrams(strs []string) [][]string {
    myMap:=make(map[string][]string)
    var res [][]string
    if len(strs)==0{
        return res
    }

    for i:=0;i<len(strs);i++{
        s1:=strings.Split(strs[i],"")
        sort.Strings(s1)
        s2:=strings.Join(s1,"")
        myMap[s2]=append(myMap[s2],strs[i])
    }
    
    for _,val:=range myMap{
        res=append(res,val)
    }
    return res
}
