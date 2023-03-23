package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/endo-checker/auth/model"
	"github.com/gookit/cache"
)

func Auth0SignIn(auth interface{}) (string, model.SignIn) {
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

	r.Header.Add("Authorization", "Bearer "+rep["access_token"].(string))

	return tkn, model.SignIn{
		Scope:     rep["scope"].(string),
		ExpiresIn: int32(rep["expires_in"].(float64)),
		IdToken:   rep["id_token"].(string),
		TokenType: rep["token_type"].(string),
	}
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
