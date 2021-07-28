package main

import (
	"fmt"
	"sort"

)

func main(){
	arr := []int{12,1234,23,555555555,24,4322,13,843}
	//arr_One := []int{1}
	kthMin,kthMax := 1,0 // Finding 2nd lowest and 2ng highest
	min,max := usingSortPackageFindKthMinAndMax(arr,kthMin,kthMax)
	fmt.Println("Minimum :",min, "Maximum : ",max)
	min,max = usingConditinalLogicFindKthMinAndMax(arr,2,1)
	fmt.Println("Minimum :",min, "Maximum : ",max)

}

func usingSortPackageFindKthMinAndMax(arr []int,minKth int,maxKth int) (int,int){
	sort.Ints(arr)
	if maxKth == 0{
		maxKth = 1
	}
	return arr[minKth],arr[(len(arr))-maxKth]
}


/* Algorithm
   1. Create Ordered map.
   2. Map each element with it's frequency.
   3.
 */
func usingConditinalLogicFindKthMinAndMax(arr []int,minKth int,maxKth int) (int,int){
	elementMap := make(map[int]int,len(arr))

	sort.Ints(arr)

	for i := range arr{
		elementMap[arr[i]] +=1
	}
	freq :=0
	for k,v := range elementMap{
		freq += v
		if freq >= minKth{
			fmt.Println("Lowest element is : ",k)
			//return k,-1
		}
	}
	elementMapp := make(map[int]int,len(arr))
	for i:= len(arr) -1  ; i >= 0  ; i--{
		fmt.Println("Arr: ",arr[i])
		elementMapp[arr[i]] +=1
	}

	fmt.Println("Arr: ",arr, "Map: ",elementMapp)
	return -1,-1

}