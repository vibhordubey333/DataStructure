package main

import (
	"fmt"
)
/*
Move Disk  1  from source  A  to destination  B
Move disk  2  from source  A  to destination  C
Move Disk  1  from source  B  to destination  C
Move disk  3  from source  A  to destination  B
Move Disk  1  from source  C  to destination  A
Move disk  2  from source  C  to destination  B
Move Disk  1  from source  A  to destination  B
Solved in no. of steps :  7
*/

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
