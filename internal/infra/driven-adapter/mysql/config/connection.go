package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	/*DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")*/

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"myuser", "12345678", "db", "3306", "order_management")
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func InitDB() {
	/*if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}*/

	var err error
	DB, err = ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")
}

func GetDB() *gorm.DB {
	return DB
}
