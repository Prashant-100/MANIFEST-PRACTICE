package main

import (
	"fmt"
	"os"
)

func main() {
	db_name := os.Getenv("DB_NAME")
	db_pass := os.Getenv("DB_PASSWORD")
	fmt.Println("DB_NAME: "+ db_name)
	fmt.Println("DB_PASSWORD: "+ db_pass)
}