// 算術演算子
package main

import "fmt"

func main() {
	fmt.Println(1 + 2)
	fmt.Println(5 - 1)
	fmt.Println(5 * 4)
	fmt.Println(60 / 3)
	fmt.Println(60 % 3)
	fmt.Println("Hello" + "World")

	n := 1
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	n += 10
	fmt.Println(n)

	n -= 5
	fmt.Println(n)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)
}
