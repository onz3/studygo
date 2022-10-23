package main

import "fmt"

func main() {
	resultado := somaTudo(1, 3, 6)
	fmt.Println(resultado)
}

func somaTudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}
