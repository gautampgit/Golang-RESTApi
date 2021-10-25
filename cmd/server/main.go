package main

import (
	"fmt"
	"log"
)

//App - struct to declare pointer
// to database connections and messageing queues
type App struct{}

//Run functions sets up the application
func (a *App) Run() error {
	fmt.Println("Setting up our App")
	return nil
}

func main() {
	fmt.Println("API v1.0")

	if err := App.Run(); err != nil {
		log.Println("Uanble to start the REST API", err)
	}
}
