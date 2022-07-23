package main

import "fmt"

var f0 int = 0
var f1 int = 1
var f2 int

func test() {
	fmt.Println("git")
}

func test2() {
	fmt.Println("git")
}

func main() {
	fmt.Println(f0, f1) //вывожу вне цикла
	for i := 3; i <= 20; i++ {
		f2 = f0 + f1
		f0 = f1
		f1 = f2
		fmt.Println(f2)
		fmt.Println(&f2)
	}

}

//+рекурсия
