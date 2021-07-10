package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "first_variable"
	//input = "FirstVariable"
	newStr := ""
	if strings.Contains(input,"_") == true {
		for i := 0 ; i < len(input) ; i++{
			if input[i] == '_'{
				i++
				newStr = newStr + strings.ToUpper(string(input[i]))
			}else{
				newStr = newStr + string(input[i])
			}
		}
	}else{
		for i := 0 ; i < len(input) ; i++{
			if unicode.IsUpper(rune(input[i])) && i == 0{
				newStr =  newStr + string(unicode.ToLower(rune(input[i])))
			}else  if unicode.IsUpper(rune(input[i])) && i != 0 {
				newStr = newStr + "_" + string(unicode.ToLower(rune(input[i])))
			}else{
				newStr = newStr + string(input[i])
			}
		}
	}
	fmt.Println("Output: ",newStr)

}
