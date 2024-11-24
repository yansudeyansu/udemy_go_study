package main

import "fmt"

//スライド
//可変長引数

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(sum(1, 2, 3))

	fmt.Println(sum(10, 20, 30, 40, 50))

	fmt.Println(sum())
}
