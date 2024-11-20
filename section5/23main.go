package main

import "fmt"

//定数
const (
	Pi       = 3.14
	URL      = "https://www.google.com"
	SiteName = "Google"
)

const (
	X = 1
	Y
	Z
	W
	V = "AAA"
	U
	A
)

const (
	c0 = iota
	c1
	c2
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)
	fmt.Println(URL)
	fmt.Println(SiteName)
	// fmt.Println(Pi)
	fmt.Println(X, Y, Z, W, V, U, A)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)

}
