package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("when's saturday? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("in two days.")

	default:
		fmt.Println("too far away.")

	}
}
