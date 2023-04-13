func maxArea(height []int) int {
    //return bruteF(height)
    res,left,right:=0,0,len(height)-1
    for left<right{
        width:=right-left
        res=max(res,min(height[left],height[right])*width)
        if height[left]<=height[right]{
            left++
        }else{
            right--
        }
    }
    return res
}

func min(a int, b int)int{
    if a<b{
        return a
    }else{
        return b
    }
}

func max(a int, b int)int{
    if a<b{
        return b
    }else{
        return a
    }
}

func bruteF(height []int) int{
    max:=0
    for i:=0;i<len(height)-1;i++{
        for j:=i+1;j<len(height);j++{
            if height[i]<=height[j]{
                max=int(math.Max(float64(height[i]*(j-i)),float64(max)))
            }else{
                max=int(math.Max(float64(height[j]*(j-i)),float64(max)))
            }
        }
    }
    return max
}
