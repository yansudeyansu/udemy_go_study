package main

import "fmt"

// 関数
// 無名関数

func main() {
	f := func(a, b int) int {
		return a + b
	}

	i := f(1, 2)
	fmt.Println(i)

	i2 := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(i2) 
}
