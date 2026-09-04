//.arrays in go are static in nature you cannot make changes once an array is initialized

package main

import (
	"fmt"
)

func main() {
	fmt.Println("This file shows the working of arrays.")

	var arr [5]string

	arr[0] = "one"
	arr[1] = "two"
	// arr[2] = "three" //if there is an empty index it shows with an extra space in the output
	arr[3] = "four"
	arr[4] = "five"

	fmt.Println("array contains:", arr)
	fmt.Println("array contains:", len(arr))

	// output:
	// This file shows the working of arrays.
	// array contains: [one two  four five].

	var arr1 = [3]int{1, 2, 3}
	fmt.Println("array contains:", arr1)
	fmt.Println("array contains:", len(arr1))
}
