package main

import "fmt"

var x = 12

func main() {
	foo()
	x=9
	fmt.Println(x)
}

func foo() {
	fmt.Println(x)
}
