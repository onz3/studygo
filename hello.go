package main

import (
	"fmt"
)

func main() {
	resultado := somaTudo(10, 20, 4, 100, 40, 45, 44, 5)
	fmt.Println(resultado)
}

func soma(a int, b int) (result int) {
	result = a + b
	return

}

func somaTudo(x ...int) int { //um array de inteiros entrando
	resultado := 0

	for _, v := range x { //faz um loop em todos os valores passados no x
		resultado += v
	}

	return resultado

}
