package main

import (
	"fmt"
)

/*
Output:
111 110 101
*/
func main() {
	one, zero, N, output := 0, 0, 3, ""
	PrintNBitBinaryNoHavingMore1sThan0s(one, zero, N, output)

}
func PrintNBitBinaryNoHavingMore1sThan0s(one, zero, N int, output string) {
	//BaseCondition
	if N == 0 {
		fmt.Print(" ", output)
		return
	}
	// 1's are freely available to us so no condition is required for them.
	output1 := output
	output1 += "1"
	PrintNBitBinaryNoHavingMore1sThan0s(one+1, zero, N-1, output1)

	/*
		O's are not freely available to us and as mentioned in problem statement their count
		should be less than equal to 1's so condition is needed.
	*/
	if one > zero {
		output2 := output
		output2 += "0"
		PrintNBitBinaryNoHavingMore1sThan0s(one, zero+1, N-1, output2)
	}

}
