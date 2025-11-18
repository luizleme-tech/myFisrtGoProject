package main

import (
	"fmt"
	"math"
	"time"
)

//func main() {
//	do(1)
//	do(2)
//	do(3)
//}
//
//func do(x int) {
//	switch x {
//	case 1:
//		fmt.Println(1)
//	case 2:
//		fmt.Println(2)
//	default:
//		fmt.Println("outra coisa")
//	}
//}

func main() {
	switch x := math.Sqrt(4; x {
	case 2:
		fmt.Println(" resultado e 2")
	default:
		fmt.Println("algo deu errado")
	}
}

//func isWeekend(x time.Time) bool {
//	switch x.Weekday() {
//	case time.Sunday, time.Saturday
//		return true
//	default:
//		return false
//	}
//}

func do( x any) {
	switch t:= x.(type) {
	case string:
		takeString(t)
	case int:
	case nil:
	}
}

func takeString(s string) {
	fmt.Println(s)
}
