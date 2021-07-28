package main

import (
	"fmt"
	"sort"

)

func main(){
	arr := []int{12,1234,23,555555555,24,4322,13,843}
	arr_One := []int{1}
	min,max := usingSortPackage(arr)
	fmt.Println("Minimum :",min, "Maximum : ",max)
	min,max = usingConditinalLogic(arr_One)
	fmt.Println("Minimum :",min, "Maximum : ",max)

}

func usingSortPackage(arr []int) (int,int){
	sort.Ints(arr)
	return arr[0],arr[len(arr)-1]
}

func usingConditinalLogic(arr []int) (int,int){
	min,max := arr[0],arr[0]
	if len(arr) <= 1{
		return arr[0],arr[0]
	}
	for i:=0 ; i < len(arr);i++{
		if arr[i] <= min{
			min = arr[i]
		}
		if arr[i] >= min{
			max = arr[i]
		}
	}

	return min,max
}