package database

import (
	"fmt"
	"github.com/kiliczsh/gondit/config"
	"github.com/kiliczsh/gondit/model"
	"strconv"

	"github.com/jinzhu/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DB, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Config("DB_HOST"), port, config.Config("DB_USER"),
			config.Config("DB_PASSWORD"), config.Config("DB_NAME")))

	if err != nil {
		panic("Gondit | Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{}, &model.Scan{})
	fmt.Println("Database Migrated")
}