package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"library-management-system/config"
	"library-management-system/handlers"
	"library-management-system/repository"
	"library-management-system/services"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	// Load environment configuration.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}

	// Connect to PostgreSQL.
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	log.Println("Connected to PostgreSQL successfully!")

	// Initialize repositories.
	bookRepo := &repository.PostgresBookRepository{DB: db}
	memberRepo := &repository.PostgresMemberRepository{DB: db}
	loanRepo := &repository.PostgresLoanRepository{DB: db}

	// Initialize the service layer.
	libraryService := services.NewLibraryService(
		bookRepo,
		memberRepo,
		loanRepo,
	)

	// Initialize API handlers.
	libraryHandler := handlers.NewLibraryHandler(libraryService, db)

	// Register API routes.
	mux := http.NewServeMux()
	libraryHandler.RegisterRoutes(mux)

	// Enable Angular frontend communication.
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4200",
		},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type", "Authorization",
		},
	}).Handler(mux)

	// Configure HTTP server.
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           corsHandler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Library Management API running on port", port)
	log.Println("API URL: http://localhost:" + port)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Fatal("Server failed: ", err)
	}
}
