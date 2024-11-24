package main

import (
	"fmt"
	"os"
)

//switch文

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i := x.(int)

	fmt.Println(i + 2)
	fmt.Printf("%T\n", i)

	// var j interface{} = 3.2
	// fmt.Println(j + 2)
	// fmt.Printf("%T\n", j)

	// f := j.(float64)
	// fmt.Println(f)
	// fmt.Printf("%T\n", f)

	switch v := x.(type) {
	case int:
		fmt.Println(v + 1)
	case float64:
		fmt.Println(v + 2)
	}

	if x == nil {
		fmt.Println("nil")
	} else if i, isInt := x.(int); isInt {
		// 型アサーションの構文解説:
		// 1. else if: 前の条件が満たされなかった場合の追加条件
		// 2. i, isInt := x.(int):
		//   - x.(int): xをint型として扱えるか確認する型アサーション
		//   - i: 変換成功時のint値を受け取る変数
		//   - isInt: 変換成功ならtrue、失敗ならfalseが入る変数
		// 3. isInt: 型アサーションが成功した場合のみ中の処理を実行
		fmt.Println(i + 2) // 型アサーション成功時、iはint型として扱える
	} else if f, isFloat64 := x.(float64); isFloat64 {
		fmt.Println(f + 2)
	}
}

//ラベルつきfor文

func main2() {
Loop:
	for {
		for {
			break Loop
		}
	}
}

// defer文
func main3() {
	defer fmt.Println("world")
	defer fmt.Println("hello")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func defer2() {
	defer main3()

	func() {
		fmt.Println("world")
	}()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// []byte は文字列をバイト列に変換します
	// Goではファイルに書き込む際は、文字列をバイトのスライス([]byte)に変換する必要があります
	// "world" という文字列を []byte 型にキャストして書き込んでいます
	file.Write([]byte("world"))
}
