package db_config

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/sarakhanx/go-auth-v1/queries"
)

func init() {
	//ให้ go สร้าง Absolute path ให้ไฟล์ .env
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)
	envPath := filepath.Join(exeDir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Println("Error loading .env file", err)
	}

	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
}

func InitDB() *pgx.Conn {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbUri := os.Getenv("DB_URI")
	// Connect to the database
	conn, err := pgx.Connect(context.Background(), dbUri)
	if err != nil {
		log.Fatal("Error initializing database", err)
	}
	//NOTE - Prepare Database Table if there a new Setup or have been reset database
	// Create the user table if it doesn't exist using the query from queries.go
	_, err = conn.Exec(context.Background(), queries.CreateUsersTable)
	if err != nil {
		log.Fatal("Error creating user table", err)
	}
	//EXPLAIN - if the connection is successfully just log status and return connection pool
	log.Println("Database connected successfully")
	return conn
}
