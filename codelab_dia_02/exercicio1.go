
package main

import (
	"fmt"
)

func main() {
	x := 50
	y := 25

	maiorNumero(x, y)
}

func maiorNumero(x int, y int) {
	fmt.Print("O maior número é: ")
	if x > y {
		fmt.Println(x)
	} else if y > x {
		fmt.Println(y)
	} else {
		fmt.Println("x e y são iguais")
	}
}

