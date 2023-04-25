package domain

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Image struct {
	ID        string     `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Title     string     `gorm:"type:VARCHAR(50);not null" valid:"required" form:"title" json:"title" example:"A Image Title"`
	Caption   string     `form:"caption" json:"caption"`
	ImageUrl  string     `gorm:"not null" valid:"required" form:"image_url" json:"image_url" example:"https://www.example.com/image.jpg"`
	UserID    string     `gorm:"type:VARCHAR(50);not null" json:"user_id"`
	User      *User      `gorm:"foreignKey:UserID;constraint:onEdit:CASCADE,onDelete:CASCADE" json:"-"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"not null;autoCreateTime" json:"updated_at,omitempty"`
	Comment   *Comment   `json:"-"`
}

func (image *Image) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(image); err != nil {
		return err
	}

	return
}

func (image *Image) BeforeEdit(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(image); err != nil {
		return err
	}
	return
}

type ImageUseCase interface {
	Get(context.Context, *[]Image) error
	Create(context.Context, *Image) error
	GetByID(context.Context, *Image, string) error
	Edit(context.Context, Image, string) (Image, error)
	Delete(context.Context, string) error
}

type ImageRepository interface {
	Get(context.Context, *[]Image) error
	Create(context.Context, *Image) error
	GetByID(context.Context, *Image, string) error
	Edit(context.Context, Image, string) (Image, error)
	Delete(context.Context, string) error
}
