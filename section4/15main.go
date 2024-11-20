package main

import "fmt"

func main() {
	var s string = "Hello World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(string(s[0]))
	fmt.Println(`test
	test
		test`)
	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[2:])
	fmt.Println(s[:])
	fmt.Println(s[0])
	fmt.Println(string(s[0]))

}
