package initializers

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "github.com/joho/godotenv"
    "os"
    "log"
)

var DB *gorm.DB

func SetupDatabase() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    user := os.Getenv("DATABASE_USER")
    password := os.Getenv("DATABASE_PASSWORD")
    dbname := os.Getenv("DATABASE_NAME")
    host := os.Getenv("DATABASE_HOST")
    port := os.Getenv("DATABASE_PORT")

    dsn := "user=" + user + " password=" + password + " dbname=" + dbname + " host=" + host + " port=" + port + " sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

}
