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

// Gets the auth0 domain and adds custom endpoints
func auth0Domain(endpoint string) (string, error) {
	dmn := os.Getenv("AUTH0_DOMAIN") + endpoint
	return dmn, nil
}

// Signs up a user with auth0
func Auth0SignUp(auth interface{}) (*model.SignUp, error) {
	json_data, err := json.Marshal(auth)
	if err != nil {
		fmt.Println(err)
	}

	url, err := auth0Domain("/dbconnections/signup")
	if err != nil {
		fmt.Println(err)
	}

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

	return rep, nil
}

// Signs in a user with auth0
func Auth0SignIn(auth interface{}) (string, *model.SignIn, error) {
	json_data, err := json.Marshal(auth)
	if err != nil {
		fmt.Println(err)
	}

	url, err := auth0Domain("/oauth/token")
	if err != nil {
		fmt.Println(err)
	}

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

	return tkn, rep, nil
}

// Gets an auth0 user from a cached token
func GetAuth0(token string) (*model.UserInfo, error) {
	tkn := getCachedTkn("token")

	url, err := auth0Domain("/userinfo")
	if err != nil {
		fmt.Println(err)
	}
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
	rep := &model.UserInfo{}
	err = json.Unmarshal(respBody, &rep)
	if err != nil {
		fmt.Println(err)
	}

	return rep, nil
}

// caches a value for a given key
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

// gets a cached value for a given key
func getCachedTkn(tokenName string) string {
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	tkn = cache.Get(tokenName)
	return tkn.(string)
}

// clears a cached value for a given key
func ClearCache(tokenName string, nullReq interface{}) string {
	cache.Register(cache.DvrFile, cache.NewFileCache(""))
	err := cache.Del(tokenName)
	if err != nil {
		fmt.Println(err)
	}
	return "Cache cleared"
}
