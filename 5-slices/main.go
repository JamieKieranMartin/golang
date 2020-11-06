package main

func main() {
	// Array
	// array length is a type / is constant
	// var arr1 [2]int
	// fmt.Println(len(arr1), cap(arr1))
	// var arr2 [3]int
	// fmt.Println(arr1, arr2)

	// Slice
	// sli1 := arr1[:]
	// sli2 := arr2[:]
	// fmt.Println(sli1, sli2)

	// Slice (pretty much python)
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(len(slice), cap(slice))
	// slice = slice[:4]
	// fmt.Println(len(slice), cap(slice))

	// Length of 0
	// slice := []int{}
	// fmt.Println(len(slice), cap(slice))

	// Set length of 0, but capacity 5
	// slice := make([]int, 0, 5)
	// fmt.Println(len(slice), cap(slice))

	// Append to Slice
	// slice := make([]int, 0)
	// fmt.Println(slice)
	// fmt.Println(len(slice), cap(slice))
	// slice = append(slice, 1)
	// fmt.Println(slice)
	// fmt.Println(len(slice), cap(slice))
}
