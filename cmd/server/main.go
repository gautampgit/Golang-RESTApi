package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gautampgit/Golang-RESTApi/internal/comment"
	"github.com/gautampgit/Golang-RESTApi/internal/database"
	"github.com/gautampgit/Golang-RESTApi/internal/transport/httphandler"
)

//App - struct to declare pointer
// to database connections and messageing queues
type App struct{}

//Run functions sets up the application
func (a *App) Run() error {
	fmt.Println("Setting up our App")

	db, err := database.NewDatabase()

	if err != nil {
		panic(err)
	}

	err = database.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	service := comment.NewService(db)

	handler := httphandler.NewHandler(*service)
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Println("Unable to start the server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("API v1.0")
	app := App{}
	if err := app.Run(); err != nil {
		log.Println("Uanble to start the REST API", err)
	}
}
