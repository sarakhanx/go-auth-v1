package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/sarakhanx/go-auth-v1/queries"
)

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

	// Create the user table if it doesn't exist using the query from queries.go
	_, err = conn.Exec(context.Background(), queries.CreateUsersTable)
	if err != nil {
		log.Fatal("Error creating user table", err)
	}
	//NOTE - if the connection is successfully just log status and return connection pool
	log.Println("Database connected successfully")
	return conn
}
