package models

type FileData struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Meta string `json:"meta"`
}
