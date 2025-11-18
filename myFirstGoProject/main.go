package main

import (
	"fmt"
	"math"
)

func main() {
	if 1 < 2 {
		fmt.Println("sim")
	}

	//if x := math.Sqrt(4); x < 10 { //shot statement
	if x := math.Sqrt(4); x < 1 { //shot statement
		fmt.Println(x)
		//} else if x > 0 {
	} else if x < 1 {
		fmt.Println(" maior que zero")
	} else {
		fmt.Println(" caiu no else")
	}
}
