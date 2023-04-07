func longestCommonPrefix(strs []string) string {
    if ( strs==nil || len(strs)==0){
        return ""
    }
    for i:=0;i<len(strs[0]);i++{
        pfx:=strs[0][i]
        for j:=1;j<len(strs);j++{
            if(i==len(strs[j]) || strs[j][i]!=pfx){
                return string(strs[0][0:i])
            }
        }
    }
    return strs[0]
}