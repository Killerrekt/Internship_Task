package model

type Response struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type LoginRes struct {
	AccessToken string `json:"access_token"`
}
