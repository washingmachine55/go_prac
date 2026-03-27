package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var logger *log.Logger

func loggerInit() {
	logger = log.NewWithOptions(os.Stderr, log.Options{
    	// ReportCaller: true,
		// ReportTimestamp: true,
		// TimeFormat: time.Second.String(),
	})
	logger.SetLevel(log.DebugLevel)

	styles := log.DefaultStyles()
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("FATAL 💀").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))
	// Add a custom style for key `err`
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true)
	
	logger.SetStyles(styles)
}

func loadEnv() (string, int, string, string, string) {
	if err := godotenv.Load("./.env"); err != nil {
		logger.Fatalf("Whoops! %v", err)
	}
	DB_HOST := os.Getenv("dbHost")
	DB_USER := os.Getenv("dbUser")
	DB_PASS := os.Getenv("dbPass")
	DB_NAME := os.Getenv("dbName")
	DB_PORT, err  := strconv.Atoi(os.Getenv("dbPort"))
	if err != nil {
		logger.Fatal(err)
	}
	return DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME 
}

func dbInit(DB_HOST string, DB_PORT int, DB_USER string, DB_PASS string, DB_NAME string) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	return db;
}

func dbCalls(db *sql.DB) {
	row, err := db.Query("SELECT VERSION()")
	if err != nil {
		logger.Fatal(err)
	}
	defer row.Close() // Important to close rows after processing

	var result string
	for row.Next() {
		// Scan copies the column values into the destination types
		err := row.Scan(&result)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("[POSTGRES] ", "Result Value", result)
	}

	// Check for any errors encountered during iteration
	if err := row.Err(); err != nil {
		logger.Fatal(err)
	}

	// fmt.Printf("%s",result)
}



func main() {
	loggerInit();
	// DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME := loadEnv();
	// db := dbInit(DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	// dbCalls(db)
	ExerciesArraysAndSlices()
	os.Exit(1)
}
