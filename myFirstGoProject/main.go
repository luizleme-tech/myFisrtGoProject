package main

import (
	"fmt"
	"myFirstGoProject/pacote"
)

func main() {
	fmt.Println("First Hello World in Go!")
	fmt.Println(pacote.Bar)
	pacote.PrintMinha()
	digaOi()
	fmt.Println(somar(1, 2))
	a, b := swap(10, 20)
	fmt.Println(a, b)
	res, rem := dividir(10, 3)
	fmt.Println(res, rem)
	f := pacote.Sum(2)
	x := f(1)
	fmt.Println(x)
	fmt.Println(pacote.Somar(10, 10, 10))
}

func digaOi() {
	fmt.Println("oi")
}

func somar(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return
}
