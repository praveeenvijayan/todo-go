package config

import (
	"fmt"
	"log"

	"github.com/praveeenvijayan/todo-go/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const host = "localhost"
const user = "postgres"
const password = "root"
const port = "5432"
const dbname = "todogo"
const sslmode = "disabled"
const TimeZone = "Asia/Kolkata"

func Connect() *gorm.DB {
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.ToDo{})
	fmt.Println("Database Connected...")

	return DB
}
