package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load() // envファイルを読み込み
	if err != nil {
		log.Fatalln(err)
	}

	if os.Getenv("GO_ENV") != "dev" {
		log.Fatalln(err)
	}
	// url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
