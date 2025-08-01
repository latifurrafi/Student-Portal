package database

import (
    "student-portal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
    godotenv.Load()

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    // ✅ migrate the Student model
    db.AutoMigrate(&models.Student{}, &models.Result{})
    
    DB = db
}
