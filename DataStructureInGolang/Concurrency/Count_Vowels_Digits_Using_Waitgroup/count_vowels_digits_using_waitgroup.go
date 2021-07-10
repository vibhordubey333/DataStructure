package main

import (
	"fmt"
	"sync"

)
/*
    input := "Vibhor Dubey 1234"
   Create 2 goroutines - countVowels & countDigits
	 output will display count of vowels & digits in input string

*/

var wg sync.WaitGroup

func main() {

	userInput := "Vibhor Dubey 1234"

	countVowels := make(chan int)
	countDigits := make(chan int)


	wg.Add(2)

	go countVowelsFunc(countVowels,userInput)
	go countDigitsFunc(countDigits,userInput)


	fmt.Println("Vowels Count: ",<-countVowels)
	fmt.Println("Digits Count: ",<-countDigits)

	wg.Wait()

}

func countVowelsFunc(input chan int, str string){

	var vowelCount int

	for i:=0 ; i < len(str)-1 ; i++{
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u'{
			vowelCount++
		}
	}
	input <- vowelCount
	wg.Done()
}

func countDigitsFunc(input chan int, str string){
	var digitCount int

	for i:=0 ; i < len(str) ; i++{
		if str[i] == '1' || str[i] == '2' || str[i] == '3' || str[i] == '4' || str[i] == '5' || str[i] == '6' || str[i] == '7' || str[i] == '8' || str[i] == '9' || str[i] == '0' {
			digitCount++
		}
	}
	input <- digitCount
	wg.Done()
}