package utils

import "time"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Image struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	ImageUrl string `json:"image_url"`
	UserID   string `json:"user_id"`
}

type GetedComment struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	ImageID   string     `json:"image_id"`
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *User      `json:"user"`
	Image     *Image     `json:"image"`
}

type ResponseDataGetedComment struct {
	Status string         `json:"status" example:"success"`
	Data   []GetedComment `json:"data"`
}

type AddComment struct {
	Message string `json:"message" example:"A comment"`
	ImageID string `json:"image_id" example:"image-123"`
}

type AddedComment struct {
	ID        string     `json:"id" example:"here is the generated comment id"`
	UserID    string     `json:"user_id" example:"here is the generated user id"`
	ImageID   string     `json:"image_id" example:"here is the generated image id"`
	Message   string     `json:"message" example:"A comment"`
	CreatedAt *time.Time `json:"created_at" example:"the created at generated here"`
}

type ResponseDataAddedComment struct {
	Status string       `json:"status" example:"success"`
	Data   AddedComment `json:"data"`
}

type EditComment struct {
	Message string `json:"message" example:"A new comment"`
}

type EditedComment struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	ImageUrl  string     `json:"image_url"`
	UserID    string     `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ResponseDataEditedComment struct {
	Status string        `json:"status" example:"success"`
	Data   EditedComment `json:"data"`
}

type ResponseMessageDeletedComment struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"your comment has been successfully deleted"`
}

type ResponseMessage struct {
	Status string `json:"status" example:"unsuccessful"`
	Data   string `json:"data" example:"the error explained here"`
}
