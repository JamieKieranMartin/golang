package main

func main() {
	// Pointers
	// Useful for performance
	// i := 12
	// p := &i
	// x := *p
	// fmt.Println(i, p, *p, x)
	// i = 15
	// fmt.Println(i, p, *p, x)

	// Struct
	type Square struct {
		W, H int
	}

	// Struct initialise
	// no nulls only nil reference to memory
	// var square Square
	// fmt.Println(square)
	// // square := Square{2, 3}
	// square := Square{
	// 	W: 2,
	// 	H: 3,
	// }
	// fmt.Println(square)
	// fmt.Println("area ==", square.H*square.W)

	// Struct method
	// func (s *Square) Area() int {
	// 	return s.H * s.W
	// }
	// fmt.Println(square.Area())

	// Pointer to struct
	// p := &Square{
	// 	H: 3,
	// 	W: 2,
	// }
}
