package main

import (
	"log"
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

	svc := &handler.SignInServer{}
	path, hndlr := pbcnn.NewAuthServiceHandler(svc)

	port := os.Getenv("PORT")
	if port == "" {
		port = addr
	}

	s := sv.Server{}
	s.ConnectServer(path, port, hndlr)
}
