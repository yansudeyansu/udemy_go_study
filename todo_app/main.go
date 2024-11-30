package main

import (
	"fmt"
	"log"

	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println("データベース接続:", models.Db)

	// データベースの接続確認
	err := models.Db.Ping()
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	fmt.Println("データベース接続成功")

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(2)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(2)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(2)
	// fmt.Println(u)

	// t := &models.Todo{}
	// t.Content = "Go言語の勉強2"
	// t.UserId = 1
	// t.CreateTodo()

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// fmt.Println(todos)

	// // ユーザーIDが1のユーザー情報を取得
	// u, _ := models.GetUser(1)
	// // そのユーザーに紐づくTodoを全て取得
	// todos, _ := u.GetTodosByUser()
	// // 取得したTodoを表示
	// fmt.Println(todos)

	// t, _ := models.GetTodo(1)
	// t.Content = "Go言語の勉強3"
	// t.UpdateTodo()
	// t, _ = models.GetTodo(1)
	// fmt.Println(t)

	controllers.StartMainServer()
	// u, err := models.GetUserByEmail("winas@winas.jp")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(u)

	// session, err := u.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// valid, err := session.CheckSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(valid)
}
