package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/gookit/cache"
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

	json_data, err := json.Marshal(auth)
	if err != nil {
		fmt.Println(err)
	}

	url := auth0Domain + "/oauth/token"
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Print(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	respBody := []byte(body)
	var rep map[string]interface{}
	json.Unmarshal(respBody, &rep)
	tkn := cacheToken(rep["access_token"].(string))

	resp := &pb.SignInResponse{
		AccessToken: tkn,
		Scope:       rep["scope"].(string),
		ExpiresIn:   int32(rep["expires_in"].(float64)),
		IdToken:     rep["id_token"].(string),
		TokenType:   rep["token_type"].(string),
	}

	r.Header.Add("Authorization", "Bearer "+rep["access_token"].(string))
	return connect.NewResponse(resp), nil
}

func (s *SignInServer) GetAccount(ctx context.Context, req *connect.Request[pb.GetAccountRequest]) (*connect.Response[pb.GetAccountResponse], error) {
	reqMsg := req.Msg.AccessToken
	token := reqMsg

	rep := getAuth0(token)

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

type UserInfo struct {
	Sub       string `json:"sub"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Picture   string `json:"picture"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"email"`
}

func cacheToken(token string) string {

	if token != "" {
		cache.Register(cache.DvrFile, cache.NewFileCache(""))
		cache.Set("token", token, cache.TwoDay)
		tkn = cache.Get("token")
	}
	return tkn.(string)
}

func getCachedTkn() string {

	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	tkn = cache.Get("token")
	return tkn.(string)
}

func getAuth0(token string) UserInfo {
	tkn := getCachedTkn()

	url := auth0Domain + "/userinfo"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("Authorization", "Bearer "+tkn)

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Print(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	respBody := []byte(body)
	rep := UserInfo{}
	json.Unmarshal(respBody, &rep)

	return rep
}
