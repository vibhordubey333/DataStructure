package main
import(
	"fmt"
	"math"
)

func main(){
	binary := []int{1,0,1,1,1,1,1,1,1,1,0,1,1,0,1,1,1,1,1}

	response := findCountOfMaxConsecutive1s(binary)
	fmt.Println("Response: ",response)
}

func findCountOfMaxConsecutive1s(binary []int) (result int){
	count := 0
	result = 0

	for _,i := range binary{
		if i == 0{
			count = 0
		}else{
			count++
			result = int(math.Max(float64(result),   float64(count)))
		}
	}
	return
}