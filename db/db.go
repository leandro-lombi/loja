package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	// A senha em uma aplicação real, não deve ser passada dessa forma
	conexao := "user=postgres dbname=leandro_loja password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
