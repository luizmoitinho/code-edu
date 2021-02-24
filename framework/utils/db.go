package utils

import (
	"log"
	"os"
	"projects/code-edu/domain"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("db.ConnectDb(): Error loading .env file - %v", err)
		return err
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Erro conecting to database db.ConnectDB()")
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db

}
