package main

import (
	"fmt"
)

func main() {
	count := 0
	TowerOfHanoi("A", "B", "C", 3, &count)
	fmt.Println("No. of steps: ", count)
}
func TowerOfHanoi(src, dest, helper string, n int, counter *int) {
	*counter++
	if n == 1 {
		fmt.Println("Move plate ", n, " from src ", src, " to destination ", dest)
		return
	}
	TowerOfHanoi(src, helper, dest, n-1, counter)
	fmt.Println("Move plate", n, "from src ", src, " to destination ", dest)
	TowerOfHanoi(helper, dest, src, n-1, counter)
}
