package main

import "fmt"

// メイン関数

// Plus は2つの整数を受け取り、その和を返す関数です
func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("noreturn")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	// i2, i3 := Div(9, 3)
	// fmt.Print(i2, i3)

	i4, _ := Div(9, 4)
	fmt.Println(i4)

	i5 := Double(100)
	fmt.Println(i5)

	Noreturn()
}
