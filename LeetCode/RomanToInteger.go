func romanToInt(s string) int {
    myMap:=make(map[rune]int)
    myMap['I']=1
    myMap['V']=5
    myMap['X']=10
    myMap['L']=50
    myMap['C']=100
    myMap['D']=500
    myMap['M']=1000
    num:=myMap[rune(s[len(s)-1])]
    for i:=len(s)-2;i>=0;i--{
        if myMap[rune(s[i])]>=myMap[rune(s[i+1])]{
            num=num+myMap[rune(s[i])]
        }else{
            num=num-myMap[rune(s[i])]
        }
    }
    return num
}