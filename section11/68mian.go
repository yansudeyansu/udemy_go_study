package main

import "fmt"

// MyInt型の定義
// int型を基にした新しい型を作成
type MyInt int

// MyInt型のメソッド
// MyInt型の値を出力するメソッド
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	// MyInt型の変数miを宣言(初期値は0)
	var mi MyInt
	fmt.Println(mi)        // 0を出力
	fmt.Printf("%T\n", mi) // main.MyIntを出力

	// int型の変数iを宣言して100を代入
	i := 100
	// MyInt型とint型は異なる型なので、足し算はできない
	// コンパイルエラーになる
	fmt.Println(mi + i)

	// MyInt型のメソッドを呼び出し
	mi.Print() // 0を出力
}
