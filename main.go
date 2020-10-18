package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/nayak-rakesh/go-project-setup/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load env vars")
	}
	routes.StartApp()
}
