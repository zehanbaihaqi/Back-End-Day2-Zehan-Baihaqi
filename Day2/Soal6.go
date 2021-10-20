package main
 

import(
	"fmt"
	"math"
)

 

func pangkat(base, pangkat int) int {

    res := math.Pow(float64(base), float64(pangkat))
	return int(res)
}

 

func main() {

   fmt.Println(pangkat(2, 3))  // 8

   fmt.Println(pangkat(5, 3))  // 125

   fmt.Println(pangkat(10, 2)) // 100

   fmt.Println(pangkat(2, 5))  // 32

   fmt.Println(pangkat(7, 3))  // 343

}