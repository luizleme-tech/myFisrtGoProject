package main

import (
	"fmt"
	"strconv"
)

var idade int

// bool
// int int8 int16 int32 int64 (inteiros positivos e negativos)
// uint uint8 uint16 uint32 uint64 uintptr (inteiros positivos)
//
// byte é a mesma coisa que uint8
//
// rune é a mesma coisa que int64
//
// float32 float64
//
// complex64 complex128
//
// string

// conversaoDireta
func main() {
	var x int = 10084
	f := float64(x)
	s := string(x) //salva em uma rune e salva no primeiro caractere da string
	sc := strconv.FormatInt(int64(x), 10)
	const i int = 100
	// somente rune, string, bool e valores numericos podem ser constantes

	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(sc)
}

/*
func main() {
	//var x bool
	var b uint8 = 10
	takeByte(b)
}

func takeByte(b byte) {
	fmt.Println(b)
}

func somar(a, b int) int {
	return a + b
}

func varsStrings() {
	//var nome, sobrenome string = "Pedro", "Pessoa"
	// var (
	// 	nome      = "Pedro"
	// 	sobrenome = "Pessoa"
	// )
	nome := "Pedro"
	sobrenome := "Pessoa"
	fmt.Println(nome, sobrenome, idade)
}*/
