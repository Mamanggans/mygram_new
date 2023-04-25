package utils

import "time"

type User struct {
	ID       string `json:"id" example:"here is the generated user id"`
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"johndoe@example.com"`
}

type SocialMedia struct {
	ID             string     `json:"id" example:"here is the generated social media id"`
	Name           string     `json:"name" example:"Example"`
	SocialMediaUrl string     `json:"social_media_url" example:"https://www.example.com/johndoe"`
	UserID         string     `json:"user_id" example:"here is the generated user id"`
	CreatedAt      *time.Time `json:"created_at" example:"here is the generated created at"`
	UpdatedAt      *time.Time `json:"updated_at" example:"here is the generated edited at"`
	User           *User      `json:"user"`
}

type SocialMedias struct {
	SocialMedias []SocialMedia `json:"social_medias"`
}

type GetedSocialMedia struct {
	SocialMedias interface{} `json:"social_medias"`
}

type ResponseDataGetedSocialMedia struct {
	Status string       `json:"status" example:"success"`
	Data   SocialMedias `json:"data"`
}

type AddSocialMedia struct {
	Name           string `json:"name" example:"Example"`
	SocialMediaUrl string `json:"social_media_url" example:"https://www.example.com/johndoe"`
}

type AddedSocialMedia struct {
	ID             string     `json:"id" example:"the social media id generated here"`
	Name           string     `json:"name" example:"Example"`
	SocialMediaUrl string     `json:"social_media_url" example:"https://www.example.com/johndoe"`
	UserID         string     `json:"user_id" example:"here is the generated user id"`
	CreatedAt      *time.Time `json:"created_at" example:"the created at generated here"`
}

type ResponseDataAddedSocialMedia struct {
	Status string           `json:"status" example:"success"`
	Data   AddedSocialMedia `json:"data"`
}

type EditSocialMedia struct {
	Name           string `json:"name" example:"New Example"`
	SocialMediaUrl string `json:"social_media_url" example:"https://www.newexample.com/johndoe"`
}

type EditedSocialMedia struct {
	ID             string     `json:"id" example:"here is the generated social media id"`
	Name           string     `json:"name" example:"New Example"`
	SocialMediaUrl string     `json:"social_media_url" example:"https://www.newexample.com/johndoe"`
	UserID         string     `json:"user_id" example:"here is the generated user id"`
	UpdatedAt      *time.Time `json:"updated_at" example:"the edited at generated here"`
}

type ResponseDataEditedSocialMedia struct {
	Status string            `json:"status" example:"success"`
	Data   EditedSocialMedia `json:"data"`
}

type ResponseMessageDeletedSocialMedia struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"your social media has been successfully deleted"`
}

type ResponseMessage struct {
	Status string `json:"status" example:"unsuccessful"`
	Data   string `json:"data" example:"the error explained here"`
}
