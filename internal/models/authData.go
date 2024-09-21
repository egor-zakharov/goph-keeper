package models

type AuthData struct {
	ID       string `json:"id,omitempty"`
	Meta     string `json:"meta"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
