package main

import (
	"fmt"
	"sync"
	"time"
)

type input struct {
	no  int
	mux sync.Mutex
}

func main() {

	in := input{no: 2}
	go in.Cube(&in)
	go in.Square(&in)
	time.Sleep(time.Second * 1)

}

func (obj *input) Square(in *input) {
	obj.mux.Lock()
	defer obj.mux.Unlock()
	obj.no = obj.no * obj.no
	fmt.Println("Square: ", obj.no)
}

func (obj *input) Cube(in *input) {
	obj.mux.Lock()
	defer obj.mux.Unlock()
	obj.no = obj.no * obj.no * obj.no
	fmt.Println("Cube: ", obj.no)
}

