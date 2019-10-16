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
}