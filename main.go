package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang-back-front/models" // Importe usando o caminho absoluto do módulo

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))
var db *sql.DB

func init() {
	var err error
	db, err = models.ConnectDB() // Corrigido para chamar ConnectDB de models
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers(db) // Corrigido para chamar GetUsers de models
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := models.GetUser(db, userID) // Corrigido para chamar GetUser de models
	if err != nil {
		log.Println(err)
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")

	err := models.CreateUser(db, &user) // Corrigido para chamar CreateUser de models
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User
	user.ID = userID
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")

	err := models.UpdateUser(db, &user) // Corrigido para chamar UpdateUser de models
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := models.DeleteUser(db, userID) // Corrigido para chamar DeleteUser de models
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
