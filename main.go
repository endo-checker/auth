package main

import (
	"net/http"
	"os"

	"github.com/endo-checker/auth/handler"
	"github.com/joho/godotenv"

	pbcnn "github.com/endo-checker/auth/internal/gen/auth/v1/authv1connect"
	sv "github.com/endo-checker/common/server"
)

var addr = ":8084"

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port != "" {
		addr = ":" + port
	}

	svc := &handler.SignInServer{}
	path, hndlr := pbcnn.NewAuthServiceHandler(svc)

	srvr := sv.Server{
		ServeMux: &http.ServeMux{},
	}

	sv.Server.ConnectServer(srvr, path, hndlr, addr)
}
