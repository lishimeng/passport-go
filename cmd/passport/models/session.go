package models

type BasicSession struct {
	Uid   string `json:"uid,omitempty"`
	Token string `json:"token,omitempty"`
}
