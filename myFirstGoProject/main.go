package main

import "fmt"

//Defer adia a chamada de uma função até que a  função principal retorne
//os argumentos sao validados imediatamente, mas a função é chamada após o retorno da função principal.
//Defer segue a oredem do Last In First Out. O escopo do Defer está sempre atrelado a função ao redor dele

//func main() {
//	x := doDefer()
//	fmt.Println(x)
//}
//
//func doDefer() int {
//	defer fmt.Println("world")
//	fmt.Println("hello")
//	return 10
//}

func main() {
	doDefer()
}

//func doDefer() {
//	defer fmt.Println(3)
//	defer fmt.Println(2)
//	fmt.Println(1)
//}

func doDefer() {
	x := 10
	defer func() {
		fmt.Println(x)
	}()

	x = 50
	fmt.Println(x)
}
