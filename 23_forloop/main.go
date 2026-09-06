package main

import "fmt"

func main() {
	fmt.Println("this file shows loops working")

	days := []string{"monday", "tuesday", "wednesday", "thurstday", "friday", "saturday", "sunday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is  %v\n", index, day)
	}

	roguevalue := 1

	for roguevalue < 10 {

		if roguevalue == 2 {
			goto lco // goto
		}

		// if roguevalue == 5 {
		// 	break
		// }
		if roguevalue == 5 {
			roguevalue++
			continue
		}
		fmt.Println("value is:", roguevalue)
		roguevalue++
	}

lco:
	fmt.Println("learnig about loops")

}
