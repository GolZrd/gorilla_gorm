package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
<<<<<<< HEAD
	conn := "user=postgres password=Egolgor23 dbname=webnotes sslmode=disable"
=======
	conn := "user=user password=Password dbname=productdb sslmode=disable"
>>>>>>> 9a3f7c0bfbe252fbacede40bbb2307e59c0caa86
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
