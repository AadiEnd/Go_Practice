package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("this shows a game with switch & case.")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice is 1 open your token")
	case 2:
		fmt.Println("move two spots")
	case 3:
		fmt.Println("move three spots")
		fallthrough
		// we get value of next case
		// value of dice is: 3
		// move three spots
		// move four spots
	case 4:
		fmt.Println("move four spots")
	case 5:
		fmt.Println("move five  spots")
	case 6:
		fmt.Println("move six spots and move again")
	default:
		fmt.Println("wrong number bhai.")
	}

}
