package domain

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        string     `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	UserID    string     `gorm:"type:VARCHAR(50);not null" json:"user_id"`
	ImageID   string     `gorm:"type:VARCHAR(50);not null" form:"image_id" json:"image_id"`
	Message   string     `gorm:"not null" valid:"required" form:"message" json:"message" example:"i am so betifull"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"not null;autoCreateTime" json:"updated_at,omitempty"`
	User      *User      `gorm:"foreignKey:UserID;constraint:opEdit:CASCADE,onDelete:CASCADE" json:"user"`
	Image     *Image     `gorm:"foreignKey:ImageID;constraint:opEdit:CASCADE,onDelete:CASCADE" json:"image"`
}

func (c *Comment) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}

	return
}

func (c *Comment) BeforeEdit(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}
	return
}

type CommentUseCase interface {
	Get(context.Context, *[]Comment, string) error
	Create(context.Context, *Comment) error
	GetByID(context.Context, *Comment, string) error
	Edit(context.Context, Comment, string) (Image, error)
	Delete(context.Context, string) error
}

type CommentRepository interface {
	Get(context.Context, *[]Comment, string) error
	Create(context.Context, *Comment) error
	GetByID(context.Context, *Comment, string) error
	Edit(context.Context, Comment, string) (Image, error)
	Delete(context.Context, string) error
}
