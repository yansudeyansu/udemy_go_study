package main

import "fmt"

// User構造体の定義
// Name: ユーザー名を格納する文字列フィールド
// Age: ユーザーの年齢を格納する整数フィールド
type User struct {
	Name string
	Age  int
	//X, Y int  // 未使用のフィールド
}

// 値渡しでUserを更新する関数
// 呼び出し元の値は変更されない
func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// ポインタ渡しでUserを更新する関数
// 呼び出し元の値も変更される
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	// 変数宣言による初期化（フィールドはゼロ値）
	var user1 User
	fmt.Println(user1) // {  0}
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	//user5 := User{30, "user5"}
	//fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)

}
