package main
import (
	"fmt"
)
func main(){

	arr := make([]int,5)
	arr = []int{1,2,3,4,5}
	fmt.Println("ReverseUsingAdditionalArray",reverseUsingAdditionalArray(arr))

}

func reverseUsingAdditionalArray(arr []int) []int{

	tArr := make([]int,len(arr))
	j := 0
	for i := len(arr)-1 ; i >= 0  ; i--{
		tArr[j] = arr[i]
		j++
	}
	return tArr
}