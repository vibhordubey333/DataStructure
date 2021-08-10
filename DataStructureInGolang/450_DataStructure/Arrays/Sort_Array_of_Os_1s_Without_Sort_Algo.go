package main
import(
	"fmt"
)
func main(){
	arr := make([]int,0)// It's not necessary to provide size as it's a slice, it will double up automatically.
	arr = []int{0,2,1,2,0,0,0,1,2}
	countingSort(arr)
}

func countingSort(arr []int){
	if len(arr) <= 1{
		fmt.Println(arr)
	}
	var(
	cnt_0 = 0
	cnt_1 = 0
	cnt_2 = 0
	)

	for i :=0 ; i < len(arr) ; i++{
		switch arr[i]{
		case 0:
			cnt_0++
			break
		case 1:
			cnt_1++
			break
		case 2:
			cnt_2++
			break
		}
	}
	i := 0

	for cnt_0 > 0{
		arr[i] = 0
		i++
		cnt_0--
	}
	for cnt_1 > 0{
		arr[i] = 1
		i++
		cnt_1--
	}
	for cnt_2 > 0{
		arr[i] = 2
		i++
		cnt_2--
	}

	fmt.Println("Output: ",arr)
}