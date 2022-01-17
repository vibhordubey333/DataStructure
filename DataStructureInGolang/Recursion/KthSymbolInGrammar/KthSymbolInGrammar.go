package main

import (
	"fmt"
	"math"
)

func main() {
	n, k := 2, 2
	fmt.Println(KthSymbolInGrammar(n, k))
}

func KthSymbolInGrammar(n, k int) int {
	if n == 1 && k == 1 {
		return 0
	}
	mid := int(math.Pow(2, float64((n-1))) / 2)
	if k <= mid {
		return KthSymbolInGrammar(n-1, k)
	} else {
		return 1 ^ KthSymbolInGrammar(n-1, k-mid) // It will give complement of first half.
	}
}
