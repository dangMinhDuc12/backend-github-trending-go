package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


type Sql struct {
	Db *sqlx.DB
	UserName string
	Password string
	Host string
	Port int
	DbName string
}

func (s *Sql) Connect() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.UserName, s.Password, s.DbName)

	s.Db = sqlx.MustConnect("postgres", connectionString)

	if err := s.Db.Ping(); err != nil {
		log.Println(err.Error())
		return
	} //short statement cho việc khai báo err và dùng err != nil trong if

	fmt.Println("Connect to db success")
}

func (s *Sql) Close() {
	s.Db.Close()
}