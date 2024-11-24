package main

import "fmt"

//スライド
//宣言,操作

func main() {
	var sl []int
	sl = append(sl, 100)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)

}
