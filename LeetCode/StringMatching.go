func strStr(haystack string, needle string) int {
    m:=len(needle)
    n:=len(haystack)
    for i:=0;i<=n-m;i++{
        for j:=0;j<m;j++{
            if needle[j]!=haystack[i+j]{
                break
            }
            if j==m-1{
                return i
            }
        }
    }
    return -1
}
