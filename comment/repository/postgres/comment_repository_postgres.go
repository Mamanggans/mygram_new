package repository

import (
	"context"
	"fmt"
	"mygram-byferdiansyah/domain"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (commentRepository *commentRepository) Get(ctx context.Context, comments *[]domain.Comment, userID string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = commentRepository.db.WithContext(ctx).Where("user_id = ?", userID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "username", "profile_image_url")
	}).Preload("Image", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "user_id", "title", "image_url", "caption")
	}).Find(&comments).Error; err != nil {
		return err
	}

	return
}

func (commentRepository *commentRepository) Create(ctx context.Context, comment *domain.Comment) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	ID, _ := gonanoid.New(16)

	comment.ID = fmt.Sprintf("your comment-%s", ID)

	if err = commentRepository.db.WithContext(ctx).Create(&comment).Error; err != nil {
		return err
	}

	return
}

func (commentRepository *commentRepository) GetByID(ctx context.Context, comment *domain.Comment, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = commentRepository.db.WithContext(ctx).First(&comment, &id).Error; err != nil {
		return err
	}

	return
}

func (commentRepository *commentRepository) Edit(ctx context.Context, comment domain.Comment, id string) (image domain.Image, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	c := domain.Comment{}

	image = domain.Image{}

	if err = commentRepository.db.WithContext(ctx).First(&c, &id).Error; err != nil {
		return image, err
	}

	if err = commentRepository.db.WithContext(ctx).Model(&c).Updates(comment).Error; err != nil {
		return image, err
	}

	if err = commentRepository.db.WithContext(ctx).First(&image, comment.ImageID).Error; err != nil {
		return image, err
	}

	return image, nil
}

func (commentRepository *commentRepository) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = commentRepository.db.WithContext(ctx).First(&domain.Comment{}, &id).Error; err != nil {
		return err
	}

	if err = commentRepository.db.WithContext(ctx).Delete(&domain.Comment{}, &id).Error; err != nil {
		return err
	}

	return
}
