package models

type MessageResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
