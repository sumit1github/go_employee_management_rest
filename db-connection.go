package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var host = "127.0.0.1"
var port = "5432"
var user = "postgres"
var password = "root"
var dbname = "go_db"

var Database *gorm.DB
var urlDSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var err error

func DataMigration() {
	Database, err = gorm.Open(postgres.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Connection failed ")
	}
	Database.AutoMigrate(&Employee{})
}
