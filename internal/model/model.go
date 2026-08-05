package model

type Chat struct{
	Chatrequest string `json:"request" binding:"required"`
}