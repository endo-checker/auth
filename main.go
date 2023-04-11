package main

import (
	"log"
	"os"

	"github.com/endo-checker/auth/handler"
	pbcnn "github.com/endo-checker/auth/internal/gen/auth/v1/authv1connect"

	sv "github.com/endo-checker/protostore/server"
	"github.com/joho/godotenv"
)

const addr = "8084"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env found: %v", err)
	}

	svc := &handler.SignInServer{}
	path, hndlr := pbcnn.NewAuthServiceHandler(svc)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = addr
	}

	// // Dapr client for downstream services
	// time.Sleep(2 * time.Second)
	// client, err := dapr.NewClient()
	// if err != nil {
	// 	log.Fatalf("failed to create dapr client: %v", err)
	// }
	// defer client.Close()

	svr := sv.Server{}
	if svr.ConnectServer(path, ":"+port, hndlr) != nil {
		log.Fatalf("Failed to connect server")
	}
}
