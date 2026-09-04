package main

import "fmt"

func swapping(x *int, y *int) (int, int) { // pass by pointers
	temp := *x
	*x = *y
	*y = temp
	fmt.Println(x, y)
	return *x, *y
}

func main() {
	fmt.Println("welcom to class on pointers.")

	// var ptr *int
	// // *string *float and all

	// fmt.Println("value of pointer is:", ptr)
	// value of empty pointer is <nil>

	first := 20
	second := 10
	fmt.Println("before swapping;", first, second)
	// var ptr = &myNumber //reference variable
	// *ptr = *ptr + 334
	// fmt.Println("value of actual pointer is:", ptr)
	// fmt.Println("value of actual pointer is:", *ptr)
	// fmt.Println("value of actual pointer is:", &myNumber)
	// fmt.Println("value of actual pointer is:", myNumber)

	swapping(&first, &second)
	fmt.Println("after swapping;", first, second)
	fmt.Println(&first, &second)

}
