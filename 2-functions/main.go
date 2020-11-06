package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, "jamie"))
	fmt.Println(divide(4, 2))

	addFunc := makeAddFunc(10)

	fmt.Println(addFunc(5))
}

func add(x, y int, z string) (int, string) {
	return x + y, "hello " + z
}

func divide(x, y int) (z int) {
	z = x / y
	return
}

func makeAddFunc(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
