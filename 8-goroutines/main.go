package main

import (
	"fmt"
	"time"
)

func doSomething() {
	count := 0
	for {
		fmt.Println(count)
		count += 1
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// go func() {
	// 	fmt.Println("HI KEVIN")
	// }()

	go doSomething()

	for {
		fmt.Println("hello")
		time.Sleep(5 * time.Second)
	}

	// 1.
	// 2.
	// 2a. go
	// 3.
	// 4.

}
