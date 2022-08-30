package main

import (
	"fmt"
	"math"
)

/*
0
01
0110
0110 1001
01101001 10010110
*/

func main() {
	//N represents no. of lines and k is the index where we've to check which digit is present from 0 and 1 on the Nth line.
	n, k := 2, 2
	fmt.Println("Digit present at ", k, " position is :", kthSymbolInGrammar(n, k))
}
/*
Algorithm:
1. Base condition: If n==1 and k ==1 then return
2. Compute mid element.
3. If k <= mid then return self_function(n-1,k)
   Else
	 return 1 ^ self_function(n-1,k-mid)
*/
func kthSymbolInGrammar(n, k int) int {
	if n == 1 && k == 1 {
		return 0
	}
	//Computing mid and then we'll compare them with k whether it is lt or gt
	mid := int(math.Pow(2, float64(n-1)) / 2)
	if k <= mid {
		return kthSymbolInGrammar(n-1, k)
	} else {
		//Returning complement of first half.
		return 1 ^ kthSymbolInGrammar(n-1, k-mid)
	}
}
