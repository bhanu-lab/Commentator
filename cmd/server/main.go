package main

import (
	"Commentator/core/database"
	"Commentator/core/transport"
	"fmt"
	"net/http"
)

type App struct{}

/*
Run - Prepares all needed for starting service
*/
func (a *App) Run() error {
	fmt.Println("Preparing Application....")

	var err error
	_, err = database.NewDataBaseConnection()
	if err != nil {
		fmt.Println("Failed creating new DB connection")
		return err
	}

	fmt.Println("[Ok]   Database connection successful")

	h := transport.NewHandler()
	h.SetupRoutes()
	fmt.Println("[Ok]   Settingup Routes successful")

	if err := http.ListenAndServe(":9090", h.Router); err != nil {
		fmt.Println("Failed to prepare application")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Hello Booting Commentator....")
	app := new(App)
	if err := app.Run(); err != nil {
		fmt.Println("Failed starting server")
		fmt.Println(err)
	}
}
