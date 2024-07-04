package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB estabelece uma conexão com o banco de dados MySQL
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:vip1234@tcp(localhost:3306)/godb")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Testar a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
