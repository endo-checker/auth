package main

import (
	"net/http"

	"github.com/endo-checker/auth/handler"

	pbcnn "github.com/endo-checker/auth/internal/gen/auth/v1/authv1connect"
	sv "github.com/endo-checker/common/server"
)

var addr = ":8084"

func main() {
	svc := &handler.SignInServer{}
	path, hndlr := pbcnn.NewAuthServiceHandler(svc)

	srvr := sv.Server{
		ServeMux: &http.ServeMux{},
	}

	sv.Server.ConnectServer(srvr, path, hndlr, addr)
}
