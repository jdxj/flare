package model

type PushInput struct {
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle,omitempty"`
	Body       string `json:"body"`
	Level      string `json:"level"`
	CipherText string `json:"ciphertext,omitempty"`
}
