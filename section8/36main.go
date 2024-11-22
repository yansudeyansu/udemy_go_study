package main

import (
	"fmt"
	"strconv"
)

//if
//条件分岐
//エラーハンドリング

func main() {
	var s string = "A"

	// strconvパッケージは文字列と他の型(数値など)との変換を行うためのパッケージです
	// Atoiは「ASCII to Integer」の略で、文字列を整数に変換する関数です
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Printf("i = %T\n", i)
}
