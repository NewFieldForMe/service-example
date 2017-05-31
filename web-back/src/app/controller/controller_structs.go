package controller

type tokenJSON struct {
	Token string `json:"token"`
}

type messageJSON struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
