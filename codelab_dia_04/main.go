package main

import "fmt"

func main() {
	array := [10]int{}
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	array[5] = 6
	array[6] = 7
	array[7] = 8
	array[8] = 9
	array[9] = 10

	fmt.Println(array)

	for i := 0; i < 10; i++ {
		fmt.Println(array[i])
	}

	values := []int{111, 1, 8, 90, 21}

	fmt.Print("Soma = ", soma(values))
}

	func soma(values []int) int {
		soma := 0
		for i := 0; i < len(values); i++ {
			soma += values[i]
		}
	return soma
}