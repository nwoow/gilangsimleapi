package models

type Requests struct {
	A float32 `json:"a" binding:"required"`
	B float32 `json:"b" binding:"required"`
}
