package models

import (
	"log"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	UUid      string
	UserId    int
	Email     string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.Id,
		&user.Uuid,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.Id,
		&user.Uuid,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}

	cmd1 := `insert into sessions (
        uuid,
        email,
        user_id,
        created_at) values (?, ?, ?, ?)`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.Id, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at from sessions
    where email = ? and user_id = ?`
	err = Db.QueryRow(cmd2, u.Email, u.Id).Scan( // 同じUUIDを使用
		&session.Id,
		&session.UUid,
		&session.Email,
		&session.UserId,
		&session.CreatedAt,
	)

	return session, err
}

func (s *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions
	where uuid = ?`
	err = Db.QueryRow(cmd, s.UUid).Scan(
		&s.Id,
		&s.UUid,
		&s.Email,
		&s.UserId,
		&s.CreatedAt,
	)
	if err != nil {
		valid = false
		return
	}
	if s.Id != 0 {
		valid = true
	}
	return valid, err
}
