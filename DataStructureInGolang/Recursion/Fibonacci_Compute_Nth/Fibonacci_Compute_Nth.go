package main
import(
	"fmt"
)
func main(){
	// Fibonacci sequence: 1 1 2 3 5 8 13
	n := 6
	output := fibonacci(n)
	fmt.Println(output)
}
func fibonacci(n int) int{
	if n <= 1 {
		return n
	}
	return fibonacci(n-1)+fibonacci(n-2)
}