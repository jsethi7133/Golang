func isValid(s string) bool {
    var stack []string
    for i:=0;i<len(s);i++{
        if (string(s[i])=="{" || string(s[i])=="[" || string(s[i])=="("){
            stack=append(stack,string(s[i]))
        }else if(string(s[i])=="}" && len(stack)!=0 && string(stack[len(stack)-1])=="{"){
                stack=stack[0:len(stack)-1]
        }else if(string(s[i])==")" && len(stack)!=0 && string(stack[len(stack)-1])=="("){
                stack=stack[0:len(stack)-1]
        }else if(string(s[i])=="]" && len(stack)!=0 && string(stack[len(stack)-1])=="["){
                stack=stack[0:len(stack)-1]                        
        }else{
            return false
        }

    }
    return len(stack)==0
}
