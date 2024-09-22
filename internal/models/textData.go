package models

type TextData struct {
	ID   string `json:"id,omitempty"`
	Meta string `json:"meta"`
	Text string `json:"text"`
}
