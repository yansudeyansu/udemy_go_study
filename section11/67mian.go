package main

import "fmt"

// User構造体の定義
// Name: ユーザー名を格納する文字列フィールド
// Age: ユーザーの年齢を格納する整数フィールド
type User struct {
	Name string
	Age  int
	//X, Y int
}

func main() {
	// マップの初期化と同時に値を設定する方法
	// キーがint型、値がUser型のマップを作成
	m := map[int]User{
		1: {Name: "user1", Age: 10}, // キー1に対してuser1を設定
		2: {Name: "user2", Age: 20}, // キー2に対してuser2を設定
	}

	// マップmの内容を出力
	fmt.Println(m)

	// User型をキー、string型を値とするマップを作成
	// ユーザー情報と居住地を紐付ける
	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}

	// マップm2の内容を出力
	fmt.Println(m2)

	// make関数を使用して空のマップを作成
	// キーがint型、値がUser型
	m3 := make(map[int]User)
	fmt.Println(m3) // 空のマップを出力

	// マップm3にキー1で新しいユーザーを追加
	m3[1] = User{Name: "user3"}
	fmt.Println(m3) // 追加後のマップを出力

	// マップmの値を順次取り出してループ処理
	// キーは使用せず、値のみを使用
	for _, v := range m {
		fmt.Println(v)
	}
}
