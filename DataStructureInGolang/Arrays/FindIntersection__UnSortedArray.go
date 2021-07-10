package main

import (
	"log"
)

func main() {
	listA := make([]int, 0)
	listB := make([]int, 0)

	listA = []int{11, 32, 0, 1, 2, 99}

	listB = []int{99, 11, 4, 2, 1, 3, 0}

	Intersection(listA, listB)
}
func Intersection(listA, listB []int) {

	m := make(map[int]bool)

	for _, v := range listA { // i depicts index and v is value
		//log.Println("i",i,":v",v)
		m[v] = true
	}

	for _, v := range listB {
		if _, ok := m[v]; ok { //_(or nok) and ok are bool , ok part will execute when this will be true
			//log.Println("nok: ",nok,": ok:",ok)
			log.Println(v)
		}

	}
}
