package models

type Article struct {
	Id       int    `json:"id"`
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft trash"`
}

func (Article) TableName() string {
	return "posts"
}
