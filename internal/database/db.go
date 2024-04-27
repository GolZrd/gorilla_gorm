package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	conn := "user=postgres password=Egolgor23 dbname=productdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Note{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
