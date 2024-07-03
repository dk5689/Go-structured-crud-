// main.go
package main

import (
	"crud_structured/controller"
	"crud_structured/repository"
	"crud_structured/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	var err error
	user := "root"
	password := ""
	host := "tcp(127.0.0.1)"
	dbName := "go_crud"
	dsn := fmt.Sprintf("%s:%s@%s/%s", user, password, host, dbName)
	//fmt.Print(dsn)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := mux.NewRouter()
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
