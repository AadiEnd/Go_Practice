package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("this is the file for slices.")

	// var fruitlist = []string{"apple", "banana", "blueberry"}
	// fmt.Printf("type of fruitslist is %T\n", fruitlist)

	// fruitlist = append(fruitlist, "guava", "mango", "peach")
	// fmt.Println(fruitlist)
	// fruitlist = append(fruitlist[1:5])
	// fmt.Println(fruitlist)
	// fruitlist = append(fruitlist[:3])
	// fmt.Println(fruitlist)

	// make syntax of slice

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 2954
	highScore[2] = 465
	highScore[3] = 486

	highScore = append(highScore, 55, 666, 776, 321)
	fmt.Println(highScore)

	copped := make([]int, len(highScore)) // copy a slice from other slice
	copy(copped, highScore)

	sort.Ints(copped)
	fmt.Println(copped)

}
