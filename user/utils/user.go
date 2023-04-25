package utils

import "time"

type RegisterUser struct {
	Age      uint   `json:"age" example:"8"`
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"secret"`
	Email    string `json:"email" example:"johndoe@example.com"`
}

type RegisteredUser struct {
	Age      uint   `json:"age" example:"8"`
	Email    string `json:"email" example:"johndoe@example.com"`
	ID       string `json:"id" example:"the user id generated here"`
	Username string `json:"username" example:"johndoe"`
}

type ResponseDataRegisteredUser struct {
	Status string         `json:"status" example:"success"`
	Data   RegisteredUser `json:"data"`
}

type LoginUser struct {
	Email    string `json:"email" example:"johndoe@example.com"`
	Password string `json:"password" example:"secret"`
}

type LoggedinUser struct {
	Token string `json:"token" example:"the token generated here"`
}

type ResponseDataLoggedinUser struct {
	Status string       `json:"status" example:"success"`
	Data   LoggedinUser `json:"data"`
}

type EditUser struct {
	Email    string `json:"email" example:"newjohndoe@example.com"`
	Username string `json:"username" example:"newjohndoe"`
}

type EditedUser struct {
	ID        string     `json:"id" example:"here is the generated user id"`
	Email     string     `json:"email" example:"newjohndoe@example.com"`
	Username  string     `json:"username" example:"newjohndoe"`
	Age       uint       `json:"age" example:"8"`
	UpdatedAt *time.Time `json:"updated_at" example:"the edited at generated here"`
}

type ResponseDataEditedUser struct {
	Status string     `json:"status" example:"success"`
	Data   EditedUser `json:"data"`
}

type ResponseMessageDeletedUser struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"your account has been successfully deleted"`
}

type ResponseMessage struct {
	Status string `json:"status" example:"unsuccessful"`
	Data   string `json:"data" example:"the error explained here"`
}
