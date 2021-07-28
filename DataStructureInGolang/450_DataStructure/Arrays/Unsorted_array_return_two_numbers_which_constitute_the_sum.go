package main

import(
	"fmt"
)
func main() {
	arr := make([]int,9)
	//arr = []int{11, 0 , -1 , 2 , -30, 1 , 7 , 2, 13}
	arr = []int{10,2,-2,-20,10}
	//sum := -29
	sum := -10
	num1, num2 := bruteForceTechnique(arr,sum)
	fmt.Println("Num1: ",num1," Num2: ",num2)
	optimizedSolution(arr,sum)
}

func bruteForceTechnique(arr []int,sum int) (int,int){

	for  _,i:= range arr{
		for _,i1 := range arr{

			if i+i1 == sum{
				fmt.Println("Two numbers to make ",sum," are :",i,", ",i1)
				return i,i1
			}
		}
	}
	return -1,-1
}


//https://www.geeksforgeeks.org/find-subarray-with-given-sum-in-array-of-integers/
func optimizedSolution(arr []int,sum int) {
	cur_sum := 0
	start := 0
	end := -1
	hashMap := make(map[int]int,len(arr))
    // 11, 0 , -1 , 2 , -30, 1 , 7 , 2, 13
	for i,v := range arr{
		cur_sum = cur_sum + v
		fmt.Println("cur_sum ",cur_sum)
		//check whether cur_sum - sum = 0, if 0 it means
		//the sub array is starting from index 0- so stop

		/*
		Say in sum we got 0 and element at 0th index is also 0 then program will
		exit mentioning no subarray exists.
		*/
		if cur_sum - sum == 0{
			start = 0
			end = i
			break
		}

		if _,ok := hashMap[cur_sum - sum] ; ok{
			start = hashMap[cur_sum-sum] + 1
			end = i
			break
		}
		hashMap[cur_sum] = i
	}
	if end == -1{
		fmt.Println("No subarray exist")

	}
	fmt.Println("Sum found between index :",start, ",",end,"\n Elements are: ")
	for i := start ; i <= end; i++{
		fmt.Print(arr[i]," ")
	}
}