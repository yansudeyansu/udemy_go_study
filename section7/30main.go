package main

import "fmt"

//関数
//関数を引数にとる

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		fmt.Println("I am a function")
	})
}
