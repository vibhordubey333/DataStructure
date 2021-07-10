package main

import (
	"fmt"
	"strings"
)

func main() {
	var newStr string
	inputString := "efiL si doog." //"Life is good"

	splitStr := strings.Split(inputString," ")

	for i:=0 ; i < len(splitStr) ; i++{
		newStr = newStr + reverse(splitStr[i]) + " "
	}
	fmt.Println("Output: ", newStr)
}


func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}