package main

import "fmt"

func switching(x, y int) (int, int) {

	temp := x
	x = y
	y = temp
	return x, y

}
func main() {
	var1 := 10
	var2 := 20
	fmt.Printf("before swap: %v, %v ", var1, var2)
	switching(var1, var2)
	fmt.Printf("after swap: %v, %v \n ", var1, var2)

}
