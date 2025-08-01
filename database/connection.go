package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"student-portal/models"
)

var DB *gorm.DB

func connect(){
	dsn := "host=localhost user=rafi password=1234 dbname=StudentPortalDb port=5432 sslmode=disable Timezone=Asia/Dhaka"
	db, err :=gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Failed To connect to the database!")
	}

	db.AutoMigrate(&models.Student{}, &models.Result{})
	DB = db
}