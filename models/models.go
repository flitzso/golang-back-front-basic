package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID    string
	Name  string
	Email string
}

// Funções de banco de dados para operações com usuários

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(db *sql.DB, id string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(db *sql.DB, user *User) error {
	_, err := db.Exec("INSERT INTO users(name, email) VALUES(?, ?)", user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *sql.DB, user *User) error {
	_, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// Função para conectar ao banco de dados MySQL
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:vip1234@tcp(localhost:3306)/godb")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Testa a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
