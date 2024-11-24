package section9

import "fmt"

func main() {
	// sl := []int{100, 200}
	// sl2 := sl

	// //スライスは参照型
	// sl2[0] = 1000

	// fmt.Println(sl)
	// fmt.Println(sl2)

	// var i int = 10
	// var p *int
	// p = &i
	// fmt.Println(p)
	// fmt.Println(*p)

	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)
}
