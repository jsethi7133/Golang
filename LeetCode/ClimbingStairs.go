func climbStairs(n int) int {
    pre:=1
    curr:=1
    for i:=1;i<n;i++{
        curr=curr+pre
        pre=curr-pre
    }
    return curr
}
