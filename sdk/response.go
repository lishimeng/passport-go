package sdk

type Response struct {
	Code    int    `json:"code,omitempty"`
	Success string `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

type AuthResponse struct {
	Response
	Token string `json:"token,omitempty"`
}

type CredentialResponse struct {
	Response
	Token string `json:"token,omitempty"`
}

// passport前端登录，选择tenant后，更换token 登录凭证（）
