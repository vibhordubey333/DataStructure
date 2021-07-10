package main

import "fmt"

func main() {
	RightShift_Approach1(5) // 5 Binary is  0000 1001
	Brian_Kerighnan_Approach2(5)
}

func RightShift_Approach1(n int) {

	count := 0
	for n > 0 {
		tempBit := 1
		fmt.Println("n & tempBit ", n&tempBit)
		if (n & tempBit) > 0 {
			count++
		}
		n = n >> 1
	}
	fmt.Println("No. of bit sets are : ", count)
}

func Brian_Kerighnan_Approach2(n int) {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
		fmt.Println("count : ", count, "N : ", n)
	}
	fmt.Println("Approach2 : No. of set bits are -> ", count)
}
