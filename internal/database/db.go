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
		sqlFile, err := os.ReadFile("internal/database/seed.sql")
		if err != nil {
			fmt.Println("failed to read sql file")
			return
		}
		db.Exec(string(sqlFile))
		fmt.Println("mock data initialized")
	}
}

func SeedData(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count == 0 {
		work1 := models.WorkItem{Name: "test work item", Description: "desc"}
		work2 := models.WorkItem{Name: "test work item 2", Description: "desc"}

		db.Create(&work1)
		db.Create(&work2)

		admin := models.User{
            Username: "admin1",
            Password: "password123", 
            Role:     models.RoleAdmin,
        }
        db.Create(&admin)

		submission := models.Submission{
            UserID:    admin.ID,
            WorkItems: []models.WorkItem{work1, work2},
        }
        db.Create(&submission)

        fmt.Println("Mock data seeded successfully.")
	}
}