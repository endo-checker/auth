package model

type SignUp struct {
	GivenName   string `json:"given_name"`
	FamilyName  string `json:"family_name"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	Connection  string `json:"connection"`
	ClientId    string `json:"client_id"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}

type UserInfo struct {
	Sub       string `json:"sub"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Picture   string `json:"picture"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"email"`
}

type SignIn struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int32  `json:"expires_in"`
	IdToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
}
