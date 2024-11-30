package models

import (
	"log"
	"time"
)

type Todo struct {
	Id        int
	Content   string
	UserId    int
	CreatedAt time.Time
}

func (t *Todo) CreateTodo() (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at	
	) values (?, ?, ?)`
	_, err = Db.Exec(cmd, t.Content, t.UserId, t.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `select * from todos where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&todo.Id,
		&todo.Content,
		&todo.UserId,
		&todo.CreatedAt,
	)
	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select * from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.Id,
			&todo.Content,
			&todo.UserId,
			&todo.CreatedAt,
		)
		// todosスライスに取得したtodoを追加する
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select * from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.Id)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.Id,
			&todo.Content,
			&todo.UserId,
			&todo.CreatedAt,
		)
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserId, t.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
