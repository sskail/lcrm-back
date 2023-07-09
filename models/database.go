package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	DbSSLMode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", DbHost, DbUser, DbPassword, DbName, DbPort, DbSSLMode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ")
	}

	err = DB.AutoMigrate(&Role{}, &User{}, &Project{}, &Comment{}, &Reply{}, &Board{}, &Task{}, &Rating{})
	if err != nil {
		log.Fatalf("Error with migrates")

	}

}
