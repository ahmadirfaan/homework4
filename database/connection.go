package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ahmadirfaan/homework4/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//ambil variable dari environment 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_database := os.Getenv("DB_DATABASE")

	//connection
	// dsn := db_user + ":" + db_pass + "tcp@(" + db_host + ":" + db_port + ")/" + db_database + 
	// "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_database)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = conn

	//migrations
	conn.AutoMigrate(
		&models.Movies{},
	)
}