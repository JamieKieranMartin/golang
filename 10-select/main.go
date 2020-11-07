package main

func main() {
	// channel := make(chan int)
	// channel1 := make(chan string)

	// go func() {
	// 	for {
	// 		channel <- rand.Int()
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		channel1 <- fmt.Sprintf("hello %d", rand.Int())
	// 		time.Sleep(5 * time.Second)
	// 	}
	// }()

	// for {
	// 	select {
	// 	case i := <-channel:
	// 		fmt.Println(i)
	// 	case s := <-channel1:
	// 		fmt.Println(s)
	// 	default:
	// 		fmt.Println("Nothing")
	// 		time.Sleep(200 * time.Millisecond)
	// 	}
	// }
}
