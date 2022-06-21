// You can edit this code!
// Click here and start typing.
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
	n, k := 2, 2
	fmt.Println("Digit present at ", k, " position is :", kthSymbolInGrammar(n, k))
}

func kthSymbolInGrammar(n, k int) int {
	if n == 1 && k == 1 {
		return 0
	}
	//Computing mid and then we'll compare them with k whether it is lt or gt
	mid := int(math.Pow(2, float64(n-1)) / 2)
	if k <= mid {
		return kthSymbolInGrammar(n-1, k)
	} else {
		return 1 ^ kthSymbolInGrammar(n-1, k-mid)
	}
}
