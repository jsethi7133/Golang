func lengthOfLongestSubstring(s string) int {
    ans:=0
    myMap:=make(map[byte]int)
    for j,i:=0,0;j<len(s);j++{
        if myMap[s[j]]!=0{
            i=int(math.Max(float64(myMap[s[j]]),float64(i)))
        }
        ans=int(math.Max(float64(ans),float64(j-i+1)))
        myMap[s[j]]=j+1
    }
    return ans
}
