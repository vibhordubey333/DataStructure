package main

import (
	"fmt"
	"sort"
)

func main() {zzzz
	inputString := "HeyThereDelilah"
	frequencyMap := make(map[string]int)

	for i := 0 ; i < len(inputString) ; i++{

		frequencyMap[string(inputString[i])] += 1


	}

	tmpStr := make([]string,0)

	for k := range frequencyMap{
		tmpStr = append(tmpStr,k)

	}

	sort.Strings(tmpStr)
	//fmt.Println(tmpStr)
	for i := 0 ; i < len(tmpStr); i ++{
		fmt.Print(tmpStr[i]+ " : ")
		fmt.Print(frequencyMap[tmpStr[i]])
		fmt.Println()
	}

}
