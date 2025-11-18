package main

import (
	"fmt"
	"sync"
)

func main() {
	var res int
	//for i := 0; i < 10; i++ {
	//	res++
	//}
	var i int
	//for ; i < 10; i++ {
	//	res++
	//}
	for i < 10 {
		res++
		i++
	}
	fmt.Println(res)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for j, elem := range arr {
	//	fmt.Println(j, elem)
	//}

	//blank identifier
	for _, elem := range arr {
		fmt.Println(elem)
	}

	for range 2 {
		fmt.Println("dentro")
	}

	const n = 10
	var wg sync.WaitGroup
	wg.Add(10)
	for k := 0; k < n; k++ {
		go func() {
			defer wg.Done()
			fmt.Println(k)
		}()
	}
	wg.Wait()
}
