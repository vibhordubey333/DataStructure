//HackerEarth Solution.
package main
import (
	"fmt"
	"strings"
)

func main() {
	var testCase int
	arrStrings := make([]string, 0)
	fmt.Scanf("%d", &testCase)
	for i := 0; i < testCase; i++ {
		temp := ""
		fmt.Scanf("%s", &temp)
		arrStrings = append(arrStrings, temp)
	}
	snakeToCamel(arrStrings, testCase)
}

func snakeToCamel(arrString []string, testCase int) {
	for k := 0; k < testCase; k++ {
		tempHolder := ""
		for i, v := range arrString[k] {

			if v >= 65 && v <= 90 { // AsCII range for A-Z.
				if i != 0 { // If index is 0th , no need to insert underscore.
					tempHolder += "_"
					tempHolder += strings.ToLower(string(v))
				} else { // Insert underscore , before character.
					tempHolder += strings.ToLower(string(v))
				}
			} else { // No need to do manipulation , insert as it is for small cap.
				tempHolder += string(v)
			}
		}
		fmt.Println(tempHolder)
	}
}
