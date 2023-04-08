func fizzBuzz(n int) []string {
    answer:=make([]string, n)
    for i:=0;i<n;i++{
        if ((i+1)%3==0 && (i+1)%5==0){
            answer[i]="FizzBuzz"
        }else if (i+1)%5==0{
            answer[i]="Buzz"
        }else if(i+1)%3==0{
            answer[i]="Fizz"
        }else{
            answer[i]=strconv.Itoa(i+1)
        }
    }
    return answer
}
