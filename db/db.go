package db

import (
	"database/sql"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=root dbname=postgres password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
