package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/service-user/domain"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/.env")

	if err != nil {
		log.Fatalf("db.ConnectDb(): Error loading .env file - %v", err)
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		fmt.Printf("Erro conecting to database db.ConnectDB(): %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db

}
