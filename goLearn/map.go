package main

import "fmt"

var f0 int = 0
var f1 int = 1
var f2 int
var i int

func main() {
	fmt.Println(f0, f1)
	for i = 3; i <= 10; i++ {
		f2 = f0 + f1
		f0 = f1
		f1 = f2
		fmt.Println(f2)
	}
}
