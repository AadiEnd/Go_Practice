package main

import "fmt"

func main() {
	//this is a way to use while loop in go
	//in go we use for loop without any init and increament clause to make it a while loop
	//loop will go until the codition becomes false

	sum := 1

	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)
}
