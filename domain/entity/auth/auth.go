package entity

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiryTime   string `json:"expires_in"`
	IDToken      string `json:"id_token"`
}

type IDPResponse struct {
	FullName string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"acct"`
}
