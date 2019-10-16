package main

import "fmt"

func main() {

	//ex1
	a := 1
	b := 2

	fmt.Println("o maior valor é: ")
	if a > b {
		fmt.Printf("b = %v", a)
	} else if a < b {
		fmt.Printf("b = %v", b)
	} else {
		fmt.Print("os valores sao iguais")
	}


	//ex2
	ano := 1999;

	if ano > (2019 - 16) {
		fmt.Printf("Nao pode votar")
	} else {
		fmt.Print("pode votar")
	}

	//ex3
	senha := "qwerty";
	entrada := "asdfg";

	if senha == entrada {
		fmt.Printf("ACESSO PERMITIDO")
	} else {
		fmt.Print("ACESSO NEGADO")
	}

	//ex4
	variavel := 1

	tipovariavel := fmt.Sprintf("%T", variavel)
	fmt.Println("tipo da variavel: ", tipovariavel)
}