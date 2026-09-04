package main

import "fmt"

func main() {
	fmt.Println("this file shows the workignon maps.")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["Go"] = "golang"
	languages["RB"] = "ruby"
	languages["py"] = "python"

	fmt.Println("list of all languages:", languages)
	fmt.Println("js is short for:", languages["JS"])

	//deletion

	delete(languages, "RB")
	fmt.Println("list of all languages:", languages)

	// loops in go lang
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)

	}
	for _, value := range languages {
		fmt.Printf("value is %v\n", value)

	}
}
