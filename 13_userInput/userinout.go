package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"

	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("give me a no. between 1-10: ")

	// comma, ok synatx

	input, _ := reader.ReadString('\n')
	fmt.Println("Your selected number: ", input)
	fmt.Printf("type of input %T", input)

}
