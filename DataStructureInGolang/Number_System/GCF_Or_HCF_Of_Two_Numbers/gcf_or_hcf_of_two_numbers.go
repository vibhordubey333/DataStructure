package main
import(
	"fmt"
)

func main(){
	num1 := 98
	num2 := 56
	output := gcd(num1,num2)
	fmt.Println("Output: ",output)
}

func gcd(num1,num2 int) int{
	if num2 == 0{
		return num1
	}
	return gcd(num2,num1%num2)
}


