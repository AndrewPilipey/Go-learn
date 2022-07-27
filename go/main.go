package main //задача на замену местами значений переменных

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x, y := 1, 2
	fmt.Println("Before x, y = ", x, y)
	swap(&x, &y)
	fmt.Println("After x, y = ", x, y)
}
