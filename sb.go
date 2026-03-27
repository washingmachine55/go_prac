package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func supabaseClient() *supabase.Client {
	if err := godotenv.Load("./.env"); err != nil {
		logger.Fatalf("Whoops! %v", err)
	}
	API_URL := os.Getenv("SB_PROJECT_URL")
	API_KEY := os.Getenv("SB_PUBLIC_KEY")
	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{})
	if err != nil {
		fmt.Println("Failed to initalize the client: ", err)
	}
	return client
}

type Tasks struct {
	TaskId    int    `json:"id"`
	TaskName  string    `json:"task_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}
