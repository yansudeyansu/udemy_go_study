package main

import "fmt"

//for
//繰り返し処理

// continue: ループの残りの部分をスキップして次のループに進む
// break: ループを抜ける
func main() {
	// 基本的なforループ
	// continueで3をスキップし、6より大きい場合はbreakでループを抜ける
	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i > 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// while文のような書き方
	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// breakを使用した単純なforループ
	// for i := 0; i < 10; i++ {
	// 	if i > 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// スライスに対するrangeループ
	// インデックスと値を取得
	// arr := []string{"A", "B", "C"}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// 配列に対するrangeループ
	// インデックスと値を取得
	// arr := [3]int{1, 2, 3}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// インデックスを無視して値のみを取得
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// スライスに対するrangeループ
	// 値のみを取得
	// sl := []string{"A", "B", "C"}
	// for _, v := range sl {
	// 	fmt.Println(v)
	// }

	// マップに対するrangeループ
	// キーと値を取得
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
