package dbexg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"gitlab.com/mahdiidarabi/exg/model"
)

func SetConnection() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("got error in loading .env")
	}
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	fmt.Println(dsn)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("got error in setConnection")
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return database, nil
}
