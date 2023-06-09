package handler

import (
	"context"
	"os"

	"github.com/bufbuild/connect-go"

	pb "github.com/endo-checker/auth/internal/gen/auth/v1"
	pbcnn "github.com/endo-checker/auth/internal/gen/auth/v1/authv1connect"
)

type SignInServer struct {
	pbcnn.UnimplementedAuthServiceHandler
}

func (s *SignInServer) CreateAccount(ctx context.Context, req *connect.Request[pb.CreateAccountRequest]) (*connect.Response[pb.CreateAccountResponse], error) {
	reqMsg := req.Msg
	auth := reqMsg.RegisterAuthUser

	auth.ClientId = os.Getenv("AUTH_CLIENT_ID")
	auth.Connection = "Username-Password-Authentication"

	rep, err := Auth0SignUp(auth)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateAccountResponse{
		Id: rep.Id,
		RegisterAuthUser: &pb.RegisterAuthUser{
			Email:      rep.Email,
			Nickname:   rep.Nickname,
			GivenName:  rep.GivenName,
			FamilyName: rep.FamilyName,
		},
	}
	if err := connect.NewResponse(resp); err != nil {
		return err, nil
	}
	return connect.NewResponse(resp), nil
}

func (s *SignInServer) SignIn(ctx context.Context, req *connect.Request[pb.SignInRequest]) (*connect.Response[pb.SignInResponse], error) {
	reqMsg := req.Msg
	auth := reqMsg.AuthUserSignIn

	auth.ClientId = os.Getenv("AUTH_CLIENT_ID")
	auth.ClientSecret = os.Getenv("AUTH_CLIENT_SECRET")
	auth.GrantType = "password"
	auth.Audience = ""

	tkn, rsp, err := Auth0SignIn(auth)
	if err != nil {
		return nil, err
	}

	resp := &pb.SignInResponse{
		AccessToken: tkn,
		IdToken:     rsp.IdToken,
		Scope:       rsp.Scope,
		ExpiresIn:   rsp.ExpiresIn,
		TokenType:   rsp.TokenType,
	}
	if err := connect.NewResponse(resp); err != nil {
		return err, nil
	}
	return connect.NewResponse(resp), nil
}

func (s *SignInServer) GetAccount(ctx context.Context, req *connect.Request[pb.GetAccountRequest]) (*connect.Response[pb.GetAccountResponse], error) {
	reqMsg := req.Msg.AccessToken
	token := reqMsg

	rsp, err := GetAuth0(token)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetAccountResponse{
		UserInfo: &pb.UserInfo{
			Sub:        rsp.Sub,
			Name:       rsp.Name,
			Nickname:   rsp.Nickname,
			Picture:    rsp.Picture,
			UpdatedAt:  rsp.UpdatedAt,
			Email:      rsp.Email,
			GivenName:  rsp.GivenName,
			FamilyName: rsp.FamilyName,
		},
	}

	if err := connect.NewResponse(resp); err != nil {
		return err, nil
	}
	return connect.NewResponse(resp), nil
}

func (s *SignInServer) SignOut(ctx context.Context, req *connect.Request[pb.SignOutRequest]) (*connect.Response[pb.SignOutResponse], error) {
	reqMsg := req.Msg
	rsp := ClearCache("token", reqMsg)

	resp := &pb.SignOutResponse{
		Message: rsp,
	}

	if err := connect.NewResponse(resp); err != nil {
		return err, nil
	}
	return connect.NewResponse(resp), nil
}
