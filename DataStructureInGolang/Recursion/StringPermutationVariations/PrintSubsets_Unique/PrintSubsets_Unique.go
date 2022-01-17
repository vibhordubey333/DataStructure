package main

import "fmt"

/*
""
b
a
ab
aa
aab

*/
var (
	hashMap = make(map[interface{}]int, 0)
)

func main() {
	printUniqueSubsets("aab", "")
	printHashMap(hashMap)

}

func printUniqueSubsets(input, output string) {
	if len(input) == 0 {
		if _, ok := hashMap[output]; !ok { //&& output != "" {
			hashMap[output] += 1
		}
		return
	}
	output1 := output
	output2 := output
	output2 += string(input[0])
	input = input[1:]
	printUniqueSubsets(input, output1)
	printUniqueSubsets(input, output2)
	return
}

func printHashMap(hashMap map[interface{}]int) {
	for k, _ := range hashMap {
		fmt.Print(" ", k)
	}
}
