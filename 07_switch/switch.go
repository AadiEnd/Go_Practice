package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Go runs on : ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS. ")
	case "linux":
		fmt.Println("Linux.")

	default:
		//windows
		fmt.Printf("%s.\n", os)
	}
	in := "hey"
	fmt.Printf("hey %T", in)
}
