package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to the time study of go lang.")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// format of time
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 "))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create time manually entering

	createDate := time.Date(2020, time.November, 19, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
