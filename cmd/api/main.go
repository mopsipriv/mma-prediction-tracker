package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"mma-prediction-tracker/internal/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, reading from system environment")
	}

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbName := getEnv("DB_NAME", "mma_tracker")

	dbPassword := mustGetEnv("DB_PASSWORD")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Connected to PostgreSQL using secure .env config")

	fighterRepo := repository.NewFighterRepository(db)

	fighters, err := fighterRepo.GetAll(ctx)
	if err != nil {
		log.Fatalf("Failed to get fighters: %v", err)
	}

	fmt.Printf("Found %d fighters in DB:\n", len(fighters))
	for _, f := range fighters {
		fmt.Printf("- [%d] %s %s (%s)\n", f.ID, f.FirstName, f.LastName, f.WeightClass)
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Critical Config Error: environment variable %s is required", key)
	}
	return value
}
