package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "MergeSort")
	arr := make([]int, 0)
	arr = []int{300, 15, 191, -6, 1, 7, 2, 4, -5}
	fmt.Println(sort(arr))
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
func sort(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	left, right := split(A)
	left = sort(left)
	right = sort(right)
	return merge(left, right)
}

func split(A []int) ([]int, []int) {
	return A[0 : len(A)/2], A[len(A)/2:]
}

func merge(A, B []int) []int {
	arr := make([]int, len(A)+len(B))
	fmt.Println("Arr:", arr)
	j, k := 0, 0

	for i := 0; i < len(arr); i++ {

		if j >= len(A) {
			arr[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			arr[i] = A[j]
			j++
			continue
		}

		if A[j] > B[k] {
			arr[i] = B[k]
			k++
		} else {
			arr[i] = A[j]
			j++
		}
	}
	return arr
}
