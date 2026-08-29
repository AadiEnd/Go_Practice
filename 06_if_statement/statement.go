package main

import (
	"fmt"
	"math"
)

func pow(v, n, lim float64) float64 {
	if v := math.Pow(v, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(2, 4, 10),
		pow(3, 3, 20),
	)
}
