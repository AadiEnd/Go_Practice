package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "exactly 10 login count"
	}

	fmt.Println(result)

	if v := 100; v < 200 {
		fmt.Println("v is low")
	}

	// if err !=nil{

	// }

}
