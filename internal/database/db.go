package database

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/l1rn/order-handler/internal/models"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(
		sqlite.Open("internal/database/orders.db"),
		&gorm.Config{},
	)
	if err != nil {
		fmt.Println("Failed to connect to db:", err)
	}

	fmt.Println("db connection established")

	err = db.AutoMigrate(
		&models.User{},
		&models.Submission{},
		&models.WorkItem{},
	)

	if err != nil {
		fmt.Println("Failed to migrate models: ", err)
	}

	fmt.Println("db migration completed")
	return db
}

func InitMockData(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		sqlFile, _ := os.ReadFile("internal/database/seed.sql")
		db.Exec(string(sqlFile))
		fmt.Println("mock data initialized")
	}
}
