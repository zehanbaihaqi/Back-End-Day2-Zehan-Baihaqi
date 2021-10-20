package main

import "fmt"


func main() {

	N := 20

	for i := 1; i < N+1 ; i++ {
		if N % i == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(N)

}