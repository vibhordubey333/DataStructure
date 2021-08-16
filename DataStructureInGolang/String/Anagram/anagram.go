package main

import (
	"fmt"
	"sort"
)

func main() {
    inputString := "abcde"
    anagramString := "ceabd"
    
    response := anagramChecker(inputString,anagramString)
    fmt.Println(inputString+":"+anagramString+" "+response)
}

func StringToRuneSlice(s string) []rune {
      var r []rune
      for _, runeValue := range s {
              r = append(r, runeValue)
      }
      return r
}

func SortStringByCharacter(s string) string {
      r := StringToRuneSlice(s)
      sort.Slice(r, func(i, j int) bool {
              return r[i] < r[j]
      })
      return string(r)
}

func anagramChecker(inputString,anagramString string) string{
    
    //Best case complexity is O(N)
    //Worst will be O(nlogn)
    
    // Return if length is not same
    if len(inputString) != len(anagramString){
        return "Not anagram"
    }
    if inputString == anagramString{
        return "Anagram"
    }
    
    inputString = SortStringByCharacter(inputString)
    anagramString = SortStringByCharacter(anagramString)
    
    if inputString != anagramString{
        return "Not anagram"
    }
    return "Are anagram"
}
