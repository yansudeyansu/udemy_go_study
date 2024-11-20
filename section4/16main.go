package main

import "fmt"

//型
// byte(uint8)型

func main() {
	// byteは予約語なので変数名として使用できません
	// 別の変数名を使用する必要があります
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))
	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

	fmt.Println(string(c[0]))
}
