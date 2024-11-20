package main

// 関数
// クロージャーの実装

import "fmt"

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("Bye"))
	fmt.Println(f("Bye1"))
	fmt.Println(f("Bye2"))
	fmt.Println(f("Bye3"))
	fmt.Println(f("Bye4"))

}
