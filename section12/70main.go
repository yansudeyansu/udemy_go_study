// interface
package main

import "fmt"

//異なる型に共通の性質を付与する

// Animalインターフェースの定義
// 動物の共通の振る舞いを定義
type Animal interface {
	Cry() string
}

// Dog構造体の定義
type Dog struct {
	Name string
}

// Dog構造体のCryメソッド
// Animalインターフェースを実装
func (d *Dog) Cry() string {
	return d.Name + ": Bow!"
}

// Cat構造体の定義
type Cat struct {
	Name string
}

// Cat構造体のCryメソッド
// Animalインターフェースを実装
func (c *Cat) Cry() string {
	return c.Name + ": Meow!"
}

func main() {
	// DogとCatのインスタンスを作成
	dog := &Dog{"Pochi"}
	cat := &Cat{"Tama"}

	// Animalインターフェース型の配列に格納
	animals := []Animal{dog, cat}

	// それぞれの動物の鳴き声を出力
	for _, animal := range animals {
		fmt.Println(animal.Cry())
	}
}
