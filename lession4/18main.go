package main

import "fmt"

// 配列型
func main() {
	var arrl [3]int
	fmt.Println(arrl)
	fmt.Printf("%T\n", arrl)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	arr5 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr5)

	fmt.Println(arr2[2])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2-1])

	arr2[2] = "300"
	fmt.Println(arr2)

	// var arr6 [4]int
	// arr6 = arr5
	// fmt.Println(arr6)

	fmt.Println(len(arrl))
}
