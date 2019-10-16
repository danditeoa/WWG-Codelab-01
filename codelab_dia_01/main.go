package main

import "fmt"

func main() {

	//ex1
	var a = ola;
	var b = pessoas;
	var c = bonitas;
	fmt.Println(a, b, c);

	//ex2
	a, b := 230, 27
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)

	soma := a + b
	subtrai := a - b
	fmt.Println("Soma: ", soma)
	fmt.Println("Subtração: ", subtrai)

	//ex3
	a, b := 230, 27
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)

	soma := a + b
	subtrai := a - b
	multiplica := (a * b)
	divide := (a/ b);
	
	fmt.Println("Soma: ", soma)
	fmt.Println("Subtração: ", subtrai)
	fmt.Println("Multipuicação: ", multiplica)
	fmt.Println("Divisão: ", divide)

	//ex4
	a, b := 230, 27
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)

	soma := a + b
	subtrai := a - b
	multiplica := (a * b)
	divide := (a/ b);
	
	fmt.Println("Soma: ", soma)
	fmt.Println("Subtração: ", subtrai)
	fmt.Println("Multipuicação: ", multiplica)
	fmt.Println("Divisão: ", divide)
	
	fmt.Printf("\n%v \n%v  \n%v  \n%v", soma, subtrai, multiplica, divide)

	//ex5
	var precoDoPao float64 = 4.60
	var precoDoAzeite float64 = 19.95
	var precoDoSuco float64 = 7
	var precoDaÁgua float64 = 2.15
	var precoDoKGMaçã float64 = 8.90
	
	/* comprar
		- 3 pães
		- 1 azeite
		- 2 garrafas de suco de laranja
		- 5 garrafas de água
		- 1.5 kg de maçã
	*/
	
	total := 3 * precoDoPao
	total += precoDoAzeite
	total += 2 * precoDoSuco
	total += 5 * precoDaÁgua
	total += 1.5 * precoDoKGMaçã
	

	fmt.Println("Total da compra: ", total)

	//ex6
	// "Var é uma palavra reservada a declaração de variaveis no go por isso não é possível declarar elas com esse nome"

	//ex7
	var X int = 15
	Y := 31
	var Z int
	Z = 47
	
	fmt.Printf("\nValores %v, %v, %v.", X, Y, Z)
	fmt.Printf("\nBase 2 %b, %b, %b.", X, Y, Z)
	fmt.Printf("\nBase 10 %d, %d, %d.", X, Y, Z)
	fmt.Printf("\nBase 16 %x, %x, %x.", X, Y, Z)

	//ex8
	// Não é possível rodar programas em go sem variaveis inicializadas
}