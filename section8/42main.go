//panicとrecover


package main

import "fmt"

func main() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	panic("runtime error")
	fmt.Println("start")
}


