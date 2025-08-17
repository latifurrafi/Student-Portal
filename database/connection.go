package database

import (
    "fmt"
    "os"
    "time"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
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

    // 👇 Disable prepared statement cache (fixes "already exists" error)
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  dsn,
        PreferSimpleProtocol: true, // disable statement caching
    }), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    // 👉 Get generic database object sql.DB to configure pooling
    sqlDB, err := db.DB()
    if err != nil {
        panic("Failed to get generic DB object: " + err.Error())
    }

    // 👉 Configure connection pool
    sqlDB.SetMaxOpenConns(100)                  // Max open connections
    sqlDB.SetMaxIdleConns(2)                   // Max idle connections
    sqlDB.SetConnMaxLifetime(5 * time.Minute)  // Reuse a connection for 5 minutes

    DB = db
    fmt.Println("✅ Database connected with connection pooling")
}
