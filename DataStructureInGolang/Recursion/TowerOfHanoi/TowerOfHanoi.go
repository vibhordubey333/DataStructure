package main

import (
	"fmt"
)

func main() {
	TowerOfHanoi("A", "B", "C", 3)
}
func TowerOfHanoi(src, dest, helper string, n int) {
	if n == 1 {
		fmt.Println("Moving plate from src ", src, " to ", dest)
		return
	}
	TowerOfHanoi(src, helper, dest, n-1)
	fmt.Println("Moving plate from src ", src, " to ", dest)
	TowerOfHanoi(helper, dest, src, n-1)
}
