package domain

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             string     `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Name           string     `gorm:"type:VARCHAR(50);not null" valid:"required" form:"name" json:"name" example:"Social Media"`
	SocialMediaUrl string     `gorm:"not null" valid:"required" form:"social_media_url" json:"social_media_url" example:"https://www.example.com/social-media"`
	UserID         string     `gorm:"type:VARCHAR(50);not null" json:"user_id"`
	CreatedAt      *time.Time `gorm:"not null;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt      *time.Time `gorm:"not null;autoCreateTime" json:"updated_at,omitempty"`
	User           *User      `gorm:"foreignKey:UserID;constraint:onEdit:CASCADE,onDelete:CASCADE" json:"user"`
}

func (s *SocialMedia) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		return err
	}

	return
}

func (s *SocialMedia) BeforeEdit(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		return err
	}
	return
}

type SocialMediaUseCase interface {
	Get(context.Context, *[]SocialMedia, string) error
	Create(context.Context, *SocialMedia) error
	GetByID(context.Context, *SocialMedia, string) error
	Edit(context.Context, SocialMedia, string) (SocialMedia, error)
	Delete(context.Context, string) error
}

type SocialMediaRepository interface {
	Get(context.Context, *[]SocialMedia, string) error
	Create(context.Context, *SocialMedia) error
	GetByID(context.Context, *SocialMedia, string) error
	Edit(context.Context, SocialMedia, string) (SocialMedia, error)
	Delete(context.Context, string) error
}
