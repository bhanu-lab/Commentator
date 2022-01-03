package main

import (
	"fmt"
	"net/http"

	"github.com/bhanu-lab/Commentator/core/comment"
	"github.com/bhanu-lab/Commentator/core/database"
	"github.com/bhanu-lab/Commentator/core/transport"
)

type App struct{}

/*
Run - Prepares all needed for starting service
*/
func (a *App) Run() error {
	fmt.Println("Preparing Application....")

	db, err := database.NewDataBaseConnection()
	if err != nil {
		fmt.Println("Failed creating new DB connection")
		return err
	}
	fmt.Println("[Ok]   Database connection successful")
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	fmt.Println("[Ok]   Databse Migration successful")

	service := comment.NewService(db)
	h := transport.NewHandler(service)
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
