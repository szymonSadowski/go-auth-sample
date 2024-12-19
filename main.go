package main

import (
	"fmt"
	"net/http"

	"github.com/goschool/crud/api"
	"github.com/goschool/crud/db"
	"github.com/goschool/crud/routes"
)

func main() {
	fmt.Println("Program starting right now")

	database, err := db.Open()

	fmt.Println("Opening database")
	if err != nil {
		panic(err)
	}

	userStore := db.NewSQLiteUserStore(database)
	userHandler := api.NewUserHandler(userStore)

	r := routes.SetupRoutes(*userHandler)
	fmt.Println("Listening on port 8081")
	http.ListenAndServe(":8081", r)

	database.Close()
	fmt.Println("Connection closed")
}
