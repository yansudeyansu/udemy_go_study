package main

import "fmt"

func main() {
	a := 10
	if a > 0 {
		fmt.Println("a is positive")
	} else if a < 0 {
		fmt.Println("a is negative")
	} else {
		fmt.Println("a is zero")
	}

	if b := 20; b > 0 {
		fmt.Println("b is positive")
	} else if b < 0 {
		fmt.Println("b is negative")
	} else {
		fmt.Println("b is zero")
	}

}
