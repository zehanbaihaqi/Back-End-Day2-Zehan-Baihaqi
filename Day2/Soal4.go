package main

import "fmt"


func main() {

	N := 20

	if 1 > N{
		for i := 2; i < N ; i++ {
		if N % i == 0 {
			fmt.Println("Bilangan Prima",i)
			break
		}else {
		fmt.Println("Bukan Bilangan ")
		}
		
	}	
	}else {
		fmt.Println("Bukan Bilangan ")
	}
	

}