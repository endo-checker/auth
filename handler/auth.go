package handler

import (
	"context"

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

const auth0Domain = "https://react-messaging.au.auth0.com"

func (s *SignInServer) SignIn(ctx context.Context, req *connect.Request[pb.SignInRequest]) (*connect.Response[pb.SignInResponse], error) {
	reqMsg := req.Msg
	auth := reqMsg.AuthUserSignIn

	godotenv.Load()
	auth.ClientId = os.Getenv("AUTH_CLIENT_ID")
	auth.ClientSecret = os.Getenv("AUTH_CLIENT_SECRET")

	tkn, rsp := Auth0SignIn(auth)

	resp := &pb.SignInResponse{
		AccessToken: tkn,
		IdToken:     rsp.IdToken,
		ExpiresIn:   rsp.ExpiresIn,
		Scope:       rsp.Scope,
		TokenType:   rsp.TokenType,
	}
	return connect.NewResponse(resp), nil
}

func (s *SignInServer) GetAccount(ctx context.Context, req *connect.Request[pb.GetAccountRequest]) (*connect.Response[pb.GetAccountResponse], error) {
	reqMsg := req.Msg.AccessToken
	token := reqMsg

	rep := GetAuth0(token)

	resp := &pb.GetAccountResponse{
		UserInfo: &pb.UserInfo{
			Sub:       rep.Sub,
			Name:      rep.Name,
			Nickname:  rep.Nickname,
			Picture:   rep.Picture,
			UpdatedAt: rep.UpdatedAt,
			Email:     rep.Email,
		},
	}
	return connect.NewResponse(resp), nil
}
