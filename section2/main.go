package main

import "fmt"

//変数
// var 変数名 型 = 値
// var := 500
var i5 int = 500

func outer() {
	var s4 string = "Golang"
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello, World!"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3) //0 ""

	i3 = 300
	s3 = "Golang"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i) //150

	//暗黙の定義
	i4 := 400
	fmt.Println(i4) //400

	i4 = 450
	fmt.Println(i4) //450

	outer()

	// fmt.Println(s4)
}
