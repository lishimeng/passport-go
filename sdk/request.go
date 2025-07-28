package sdk

// PasswordRequest 密码式
type PasswordRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type CredentialRequest struct {
	AppId  string `json:"appId,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type CodeRequest struct {
	Code string `json:"code"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RevokeRequest struct {
	RefreshToken string `json:"refresh_token"`
}
