package main

import "fmt"

func ehPrimo(numero int) bool {
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numeros := []int{3, 5, 6, 7, 10, 11, 22, 32, 43, 111}
	for _, numero := range numeros {
		if ehPrimo(numero) {
			fmt.Printf("%d é primo.\n", numero)
		}
	}
}
