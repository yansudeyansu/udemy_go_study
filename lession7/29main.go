package main

import "fmt"

// 関数

func ReturnFunc() func() {
	return func() {
		fmt.Println("I am anonymous function")
	}
}

func main() {
	f := ReturnFunc()
	f()

}
