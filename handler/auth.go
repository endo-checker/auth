package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/joho/godotenv"

	pb "github.com/endo-checker/auth/internal/gen/auth/v1"
	pbcnn "github.com/endo-checker/auth/internal/gen/auth/v1/authv1connect"
)

type SignInServer struct {
	pbcnn.UnimplementedAuthServiceHandler
}

var tkn interface{}

func (s *SignInServer) CreateAccount(ctx context.Context, req *connect.Request[pb.CreateAccountRequest]) (*connect.Response[pb.CreateAccountResponse], error) {
	reqMsg := req.Msg
	auth := reqMsg.RegisterAuthUser

	godotenv.Load()
	auth.ClientId = os.Getenv("AUTH_CLIENT_ID")
	rep := Auth0SignUp(auth)

	resp := &pb.CreateAccountResponse{
		RegisterAuthUser: &pb.RegisterAuthUser{
			Email:      rep.Email,
			Nickname:   rep.Nickname,
			GivenName:  rep.GivenName,
			FamilyName: rep.FamilyName,
		},
	}
	return connect.NewResponse(resp), nil
}

func (s *SignInServer) SignIn(ctx context.Context, req *connect.Request[pb.SignInRequest]) (*connect.Response[pb.SignInResponse], error) {
	reqMsg := req.Msg
	auth := reqMsg.AuthUserSignIn

	fmt.Println(auth)

	godotenv.Load()
	auth.ClientId = os.Getenv("AUTH_CLIENT_ID")
	auth.ClientSecret = os.Getenv("AUTH_CLIENT_SECRET")

	tkn, rsp := Auth0SignIn(auth)

	resp := &pb.SignInResponse{
		AccessToken: tkn,
		IdToken:     rsp.IdToken,
		Scope:       rsp.Scope,
		ExpiresIn:   rsp.ExpiresIn,
		TokenType:   rsp.TokenType,
	}
	return connect.NewResponse(resp), nil
}

func (s *SignInServer) GetAccount(ctx context.Context, req *connect.Request[pb.GetAccountRequest]) (*connect.Response[pb.GetAccountResponse], error) {
	reqMsg := req.Msg.AccessToken
	token := reqMsg

	rsp := GetAuth0(token)

	resp := &pb.GetAccountResponse{
		UserInfo: &pb.UserInfo{
			Sub:       rsp.Sub,
			Name:      rsp.Name,
			Nickname:  rsp.Nickname,
			Picture:   rsp.Picture,
			UpdatedAt: rsp.UpdatedAt,
			Email:     rsp.Email,
		},
	}
	return connect.NewResponse(resp), nil
}

func (s *SignInServer) SignOut(ctx context.Context, req *connect.Request[pb.SignOutRequest]) (*connect.Response[pb.SignOutResponse], error) {
	reqMsg := req.Msg
	rsp := ClearCache(reqMsg)

	resp := &pb.SignOutResponse{
		Message: rsp,
	}
	return connect.NewResponse(resp), nil
}
