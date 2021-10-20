package main

import "fmt"

func main() {
	N := 80
	
	if N >= 80 && N <=100 {
		fmt.Println("Nilai anda A = ",N)
	} else if N >= 65 && N <= 79 {
		fmt.Println("Nilai anda B = ",N)
	} else if N >= 50 && N <= 64 {
		fmt.Println("Nilai anda C = ",N)
	} else if N >= 35 && N <= 49 {
		fmt.Println("Nilai anda D = ",N)
	}else if N >= 0 && N <= 34 {
		fmt.Println("Nilai anda D = ",N)
	}else {
		fmt.Println("Nilai Invalid")
	}

	



}