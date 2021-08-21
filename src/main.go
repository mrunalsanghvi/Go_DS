package main

import (
	"fmt"

	"github.com/mrunalsanghvi/Go_DS/pkg/hello"
)

func main() {
	fmt.Println(hello.Helloworld())
	a := []int{3, 6}
	fmt.Printf("%v, %T, len:%v\n", a, a, len(a))
}
