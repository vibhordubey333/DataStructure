package main
import(
	"fmt"
	"math"
)

func main(){
	output := findNthPrimeNumber(5)
	fmt.Println("Output:  ",output)
}
func findNthPrimeNumber(num int) int{

	tmpNum := 0
	latest := 0
	for tmpNum != num{

		count := 0
		for i:=2 ; i <= int(math.Sqrt(float64(num)));i++{
			if num % i == 0 {
				count++
				break
			}
		}
		if count == 0 {
			tmpNum++
			latest = num
		}
		num++
	}
	return latest

}