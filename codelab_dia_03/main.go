package main

import "fmt"

func main() {
	
	//ex1
	for i := 1; i <= 100; i++ {
		if (i % 3 == 0) {
			fmt.Println(i)
		}
	}

	//ex2
	for i := 1; i <= 100; i++ {
		if (i % 9 == 0) {
			fmt.Println(i)
		}
	}

	//ex3
	var teste = 3
	if (teste % 5 == 0) {
		fmt.Printf("O numero é multiplo de 5")
	} else {
		fmt.Printf("O numero não é multiplo de 5")
	}
}