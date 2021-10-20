package main

import (
	"fmt"
//	"math"
)

func main() {
	//pi := math.Pi
	pi := 3.14
	T := 20.0
	r := 4.0

	tabung := 2 * pi * r * (T + r)
	

	fmt.Println(tabung)

}