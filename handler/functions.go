package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/endo-checker/auth/model"
	"github.com/gookit/cache"
)

const auth0Domain = "https://react-messaging.au.auth0.com"

func Auth0SignIn(auth interface{}) (string, *model.SignIn) {
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
	json.Unmarshal(respBody, rep)
	tkn := cacheToken(rep.AccessToken, rep.ExpiresIn)

	return tkn, rep
}

func GetAuth0(token string) model.UserInfo {
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
	rep := model.UserInfo{}
	json.Unmarshal(respBody, &rep)

	return rep
}

func cacheToken(token string, expires int32) string {
	if token != "" {
		cache.Register(cache.DvrFile, cache.NewFileCache(""))
		cache.Set("token", token, time.Duration(expires)*time.Second)
		tkn = cache.Get("token")
	}
	return tkn.(string)
}

func getCachedTkn() string {
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	tkn = cache.Get("token")
	return tkn.(string)
}

func ClearCache(nullReq interface{}) string {
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	cache.Del("token")

	return "Cache cleared"
}
