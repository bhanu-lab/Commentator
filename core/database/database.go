package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

/*
NewDataBaseConnection - creates New databse connection with gorm
returns db,error if any
*/
func NewDataBaseConnection() (*gorm.DB, error) {
	username := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	connectstring := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, database, pass)

	fmt.Printf("Received connection details : %s \n", connectstring)

	db, err := gorm.Open("postgres", connectstring)
	if err != nil {
		return db, err
	}

	if err = db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil
}
