package main

import "fmt"

// 関数
// ジェネレーター
func IntGeners() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := IntGeners()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}
