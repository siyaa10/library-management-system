package config

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDatabase establishes a connection pool with PostgreSQL.
func ConnectDatabase() (*pgxpool.Pool, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	// Build a safe PostgreSQL connection URL.
	dbURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   net.JoinHostPort(host, port),
		Path:   dbName,
	}

	query := dbURL.Query()
	query.Set("sslmode", sslMode)
	dbURL.RawQuery = query.Encode()

	// Create a connection pool.
	pool, err := pgxpool.New(context.Background(), dbURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	// Verify the database connection.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	return pool, nil
}
