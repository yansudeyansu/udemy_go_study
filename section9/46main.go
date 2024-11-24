package main

import "fmt"

//スライス
//append,make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	for _, v := range sl {
		fmt.Println(v)
	}

	s2 := []string{"A", "B", "C"}
	for _, v := range s2 {
		fmt.Println(v)
	}

	fmt.Println(sl[0])
	fmt.Println(sl[1])
	fmt.Println(sl[2])

	fmt.Println(sl[0:2])

	fmt.Println(cap(sl))
}
