package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ResponseID struct {
	ID primitive.ObjectID `json:"id"`
}

type ResponseData struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseList struct {
	Data    any    `json:"data"`
	Total   int64  `json:"total"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
