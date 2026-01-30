package database

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/l1rn/order-handler/internal/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(
		sqlite.Open("internal/database/orders.db"),
		&gorm.Config{},
	)
	if err != nil {
		fmt.Println("Failed to connect to db:", err)
	}

	fmt.Println("db connection established")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Submission{},
		&models.WorkItem{},
	)

	if err != nil {
		fmt.Println("Failed to migrate models: ", err)
	}

	fmt.Println("db migration completed")
}

func InitMockData() {
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		sqlFile, _ := os.ReadFile("internal/database/seed.sql")
		DB.Exec(string(sqlFile))
		fmt.Println("mock data initialized")
	}
}
