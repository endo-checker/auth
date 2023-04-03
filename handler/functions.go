package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/endo-checker/auth/model"
	"github.com/gookit/cache"
)

var tkn interface{}

func Auth0SignUp(auth interface{}) *model.SignUp {
	auth0Domain := os.Getenv("AUTH0_DOMAIN")

	json_data, err := json.Marshal(auth)
	if err != nil {
		fmt.Println(err)
	}

	url := auth0Domain + "/dbconnections/signup"
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
	rep := &model.SignUp{}
	err = json.Unmarshal(respBody, rep)
	if err != nil {
		fmt.Println(err)
	}

	return rep
}

func Auth0SignIn(auth interface{}) (string, *model.SignIn) {
	auth0Domain := os.Getenv("AUTH0_DOMAIN")

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
	rep := &model.SignIn{}
	err = json.Unmarshal(respBody, rep)
	if err != nil {
		fmt.Println(err)
	}
	tkn := cacheToken("token", rep.AccessToken, rep.ExpiresIn)

	return tkn, rep
}

func GetAuth0(token string) model.UserInfo {
	auth0Domain := os.Getenv("AUTH0_DOMAIN")

	tkn := getCachedTkn("token")
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
	rep := model.UserInfo{}
	err = json.Unmarshal(respBody, &rep)
	if err != nil {
		fmt.Println(err)
	}

	return rep
}

func cacheToken(tokenName, token string, expires int32) string {
	if token != "" {
		cache.Register(cache.DvrFile, cache.NewFileCache(""))
		err := cache.Set(tokenName, token, time.Duration(expires)*time.Second)
		if err != nil {
			fmt.Println(err)
		}
		tkn = cache.Get(tokenName)
	}
	return tkn.(string)
}

func getCachedTkn(tokenName string) string {
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	tkn = cache.Get(tokenName)
	return tkn.(string)
}

func ClearCache(tokenName string, nullReq interface{}) string {
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	err := cache.Del(tokenName)
	if err != nil {
		fmt.Println(err)
	}

	return "Cache cleared"
}
