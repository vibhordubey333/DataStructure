package main

import (
	"fmt"
)

const MAXCHAR uint = 50

func main() {
	str := "AbracaDabra"
	srchStr := "ra"
	prime := uint(13)
	rabinKarp(str, srchStr, prime)
}

func rabinKarp(str, srchStr string, prime uint) {
	strLen := len(str)
	srchStrLen := len(srchStr)
	strHash := uint(0)
	srchStrHash := uint(0)
	hash := uint(1)

	for i := 0; i < srchStrLen; i++ {
		hash = (hash * MAXCHAR) % prime
		fmt.Println("Hash: ", hash)
	}

	//Calculate hash value for string and searchString.
	for i := 0; i < srchStrLen-1; i++ {
		strHash = (MAXCHAR*strHash + uint(str[i])) % prime
		srchStrHash = (MAXCHAR*srchStrHash + uint(srchStr[i])) % prime
		fmt.Println("srchStrHash", srchStrHash, "strHash: ", strHash)
	}

	for i := 0; i <= (strLen - srchStrLen); i++ {
		if strHash == srchStrHash { // If hash matches.
			fmt.Println("Found match")
			for j := 0; j < srchStrLen; j++ {
				if str[i+j] != srchStr[j] {
					break
				}

				if j == srchStrLen {
					fmt.Println("Pattern found at : %d", i)
				}
			}
		}

		if i < (strLen - srchStrLen) {
			strHash = (MAXCHAR*(strHash-uint(str[i])*hash) + uint(str[i+srchStrLen])) % prime
			if strHash < 0 {
				strHash = strHash + MAXCHAR
			}
			//fmt.Println("strHash Again:", strHash)
		}
	}

}
