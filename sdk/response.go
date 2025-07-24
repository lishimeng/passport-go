package sdk

type Response struct {
	Code    int    `json:"code,omitempty"`
	Success string `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

type OAuthResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	// AccessToken 的有效时间 seconds
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

type _oauthResponse struct {
	Response
	OAuthResponse
}

type CredentialResponse struct {
	Response
	Token string `json:"token,omitempty"`
}

// passport前端登录，选择tenant后，更换token （登录凭证），并且需要旧token失效
