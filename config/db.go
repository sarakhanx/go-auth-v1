package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/sarakhanx/go-auth-v1/queries" // Import the queries package
)

func InitDB() *pgx.Conn {
	// Connect to the database
	conn, err := pgx.Connect(context.Background(), "postgresql://admin:superUser@localhost:5432/postgres")
	if err != nil {
		log.Fatal("Error initializing database", err)
	}

	// Create the user table if it doesn't exist using the query from queries.go
	_, err = conn.Exec(context.Background(), queries.CreateUsersTable)
	if err != nil {
		log.Fatal("Error creating user table", err)
	}

	return conn
}
