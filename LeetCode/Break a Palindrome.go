func breakPalindrome(palindrome string) string {
    if len(palindrome)<2{
        return ""
    }
    for i:=0;i<len(palindrome)/2;i++{
        if string(palindrome[i])!="a"{
            return strings.Replace(palindrome,string(palindrome[i]),"a",1)
        }
    }
    return palindrome[:len(palindrome)-1]+"b"
}
