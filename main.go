package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/endo-checker/auth/handler"
	pbcnn "github.com/endo-checker/auth/internal/gen/auth/v1/authv1connect"
	sv "github.com/endo-checker/protostore/server"
	"github.com/joho/godotenv"
)

var addr = ":8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env found: %v", err)
	}

	// Create a new server
	svc := &handler.SignInServer{}
	path, hndlr := pbcnn.NewAuthServiceHandler(svc)

	srvr := sv.Server{
		ServeMux: &http.ServeMux{},
	}

	// start server
	port := os.Getenv("PORT")
	if port != "" {
		addr = ":" + port
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := srvr.ConnectServer(ctx, path, hndlr, addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
