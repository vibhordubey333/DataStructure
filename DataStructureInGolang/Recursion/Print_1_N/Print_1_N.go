package main
import(
"fmt"
)
func main(){
    n := 7
    Print1ToN(n)
}

func Print1ToN(n int){
    if n == 0{               //Base Condition.
        return
    }
    Print1ToN(n-1)          //Hypothesis, how our function will behave.
    fmt.Println(n)          //Induction step.
}