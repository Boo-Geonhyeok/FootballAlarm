package database

import (
	"fmt"
	"os"

	models "github.com/Boo-Geonhyeok/football-alarm/Models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	// dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPwd, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Fixture{})
	DB.AutoMigrate(&models.Prediction{})
}
