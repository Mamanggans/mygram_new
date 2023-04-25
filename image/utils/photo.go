package utils

import (
	"time"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type GetedImage struct {
	ID        string     `json:"id"`
	Title     string     `json:"title,"`
	Caption   string     `json:"caption"`
	ImageUrl  string     `json:"image_url"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *User      `json:"user"`
}

type ResponseDataGetedImage struct {
	Status string       `json:"status" example:"success"`
	Data   []GetedImage `json:"data"`
}

type AddImage struct {
	Title    string `json:"title" example:"A Title"`
	Caption  string `json:"caption" example:"A caption"`
	ImageUrl string `json:"image_url" example:"https://www.example.com/image.jpg"`
}

type AddedImage struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	ImageUrl  string     `json:"image_url"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type ResponseDataAddedImage struct {
	Status string     `json:"status" example:"success"`
	Data   AddedImage `json:"data"`
}

type EditImage struct {
	Title    string `json:"title" example:"A new title"`
	Caption  string `json:"caption" example:"A new caption"`
	ImageUrl string `json:"image_url" example:"https://www.example.com/new-image.jpg"`
}

type EditedImage struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	ImageUrl  string     `json:"image_url"`
	UserID    string     `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ResponseDataEditedImage struct {
	Status string      `json:"status" example:"success"`
	Data   EditedImage `json:"data"`
}

type ResponseMessageDeletedImage struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"your image has been successfully deleted"`
}

type ResponseMessage struct {
	Status string `json:"status" example:"unsuccessful"`
	Data   string `json:"data" example:"the error explained here"`
}
