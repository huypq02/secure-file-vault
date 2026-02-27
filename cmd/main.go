package main

import (
	"fmt"
	"log"

	"github.com/huypq02/secure-file-vault/internal/bootstrap"
)

func main() {
	fmt.Println("Secure File Vault - Starting...")

	// Initialize application using Wire-generated dependency injection
	app, err := bootstrap.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start the application (router + scheduler)
	fmt.Println("Server starting on :8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
