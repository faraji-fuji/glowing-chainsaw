package main

import (
	"fmt"
	"log"

	"github.com/faraji-fuji/src/routes"
	"github.com/joho/godotenv"
	"github.com/ochom/gutils/helpers"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	port := ":" + helpers.GetEnv("PORT", "8080")
	fmt.Print("Server running on ", port)

	r := routes.Init()
	if err := r.Run(port); err != nil {
		fmt.Printf("Failed to start server: %s", err.Error())
	}
}
