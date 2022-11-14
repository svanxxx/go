package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type User struct {
	id    uint64
	name  string
	pass  string
	email string
}

// interface definition
type IUserManager interface {
	FindUser(name string) *User
	RegisterUser(name string, password string, email string) (*User, error)
	Login(name string, password string) (*User, error)
	Connect()
	Disconnect()
	Connected() bool
}

type UserManager struct {
	db *sql.DB
}

func (man *UserManager) Disconnect() {
	if man.Connected() {
		man.db.Close()
		man.db = nil
	}
}
func (man *UserManager) Connected() bool {
	return man.db != nil
}

func (man *UserManager) Connect() bool {
	man.Disconnect()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable binary_parameters=yes", "127.0.0.1", 5433, "go", "go", "users")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return false
	}
	man.db = db
	return true
}

func (man *UserManager) FindUser(name string) (*User, error) {
	if !man.Connected() {
		if !man.Connect() {
			return nil, errors.New("Database not connected")
		}
	}
	var user User
	var err = man.db.QueryRow("select * from users u where u.name = $1", name).Scan(&user.id, &user.name, &user.pass, &user.email)
	if err != nil {
		return nil, nil
	}
	return &user, nil
}

func (man *UserManager) Login(name string, password string) (*User, error) {
	user, _ := man.FindUser(name)
	if user == nil {
		return nil, errors.New("User not found")
	}
	if GetMD5Hash(password) != user.pass {
		return nil, errors.New("Invalid password")
	}
	return user, nil
}

func (man *UserManager) RegisterUser(name string, password string, email string) (*User, error) {
	if !man.Connected() {
		if !man.Connect() {
			return nil, errors.New("Database not connected")
		}
	}
	if len(name) < 1 || len(password) < 1 || len(email) < 1 {
		return nil, errors.New("Name, password and email cannot be empty")
	}
	user, _ := man.FindUser(name)
	if user != nil {
		return nil, errors.New("user already exists")
	}
	_, err := man.db.Exec("insert into users (name, pass, email) values($1,$2,$3)", name, GetMD5Hash(password), email)
	if err != nil {
		panic(err)
	}
	user, _ = man.FindUser(name)
	return user, nil
}
