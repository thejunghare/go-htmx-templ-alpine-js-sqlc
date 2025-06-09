package db

import (
	"fmt"

	"github.com/thejunghare/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=1010 dbname=Tasks port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connecting port 5432!")
	}
	DB = database
	DB.AutoMigrate(&models.Tasks{})
}
