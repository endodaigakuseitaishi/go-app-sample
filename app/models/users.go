package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos []Todo
}

type Session struct {
	ID int
	UUID string
	Email string
	UserId int
	CreatedAt time.Time
}


func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

		fmt.Println(cmd)
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
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID, 
		&user.UUID, 
		&user.Name, 
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ログインページにて入力されたemail情報を用いてsession情報を取得する
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	return user, err
}

// sessionを作成
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid, 
		email, 
		user_id, 
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at
	 from sessions where user_id = ? and email = ?`

	 err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID, 
		&session.UUID, 
		&session.Email, 
		&session.UserId, 
		&session.CreatedAt)

		return session, err
}

// sessionがDBに存在するか否か
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID, 
		&sess.UUID, 
		&sess.Email, 
		&sess.UserId, 
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}

	// 初期値出ない時
	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at from users where id = ?`
	err = Db.QueryRow(cmd, sess.ID).Scan(
		&user.ID, 
		&user.UUID, 
		&user.Name, 
		&user.Email, 
		&user.CreatedAt)
	
	return user, err
}
