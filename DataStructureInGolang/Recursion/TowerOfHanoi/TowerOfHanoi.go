package main

import (
	"fmt"
)

func main() {
	TowerOfHanoi("A", "B", "C", 3)
}
func TowerOfHanoi(src, dest, helper string, n int) {
	if n == 1 {
		fmt.Println("Move plate ", n, " from src ", src, " to destination ", dest)
		return
	}
	TowerOfHanoi(src, helper, dest, n-1)
	fmt.Println("Move plate", n, "from src ", src, " to destination ", dest)
	TowerOfHanoi(helper, dest, src, n-1)
}
