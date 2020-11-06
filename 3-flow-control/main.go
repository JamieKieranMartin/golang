package main

func main() {
	// Simple for loop
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum++
	// }
	// fmt.Println(sum)

	// No while loop
	// not this >> while {}
	// count := 10
	// for count > 0 {
	// 	count--
	// 	fmt.Println(count)
	// }

	// Infinite loop
	// for i := 0; ; {
	// 	fmt.Println(i)
	// }

	// Break and Continue
	// for i := 0; ; i++ {
	// 	if i == 10 {
	// 		break
	// 	}

	// 	if i == 3 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }

	// If
	// s := "hello"
	// if s == "hello" {
	// 	fmt.Println(s)
	// } else if s == "world" {
	// } else {
	// 	fmt.Println("world")
	// }
	// if s == "hello" {
	// 	fmt.Println(s)
	// } else {
	// 	fmt.Println("world")
	// }
	// if s == "hello" {
	// 	fmt.Println(s)
	// 	return
	// }
	// fmt.Println("world")

	// If short statement
	// if v := 2 + 2; v == 3 {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("Not equal!", "expected", v)
	// }
	// v doesn't exist
	// fmt.Println(v)

	// Switch statement
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println(i, "equal to", 1)
	// case 2:
	// 	fmt.Println(i, "equal to", 2)
	// default:
	// 	fmt.Println(i, "equal to something unexpected!")
	// }

	// Switch if else cleaner
	// i := 0
	// a := 1
	// switch {
	// case a == 0 && i == 1:
	// case i > 2:
	// case a == 1:
	// }

	// Defer
	// defer fmt.Println("world")
	// fmt.Println("hello")

	// Defer function
	// deferFunc()
	// fmt.Println("kevin")
	// func deferFunc() {
	// 	defer fmt.Println("world")
	// 	fmt.Println("hello")
	// }

	// Defer loop
	// fmt.Println("starting")
	// defer fmt.Println("done done")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("loop start", i)
	// 	defer fmt.Println(i)
	// 	fmt.Println("loop end", i)
	// }
	// fmt.Println("done")

}
