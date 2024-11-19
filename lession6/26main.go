package main

// 論理演算子

import "fmt"

func main() {
	fmt.Println(true && false == true)
	fmt.Println(true || false == true)
	fmt.Println(!true == false)

	fmt.Println(!true)
	fmt.Println(!false)
}
