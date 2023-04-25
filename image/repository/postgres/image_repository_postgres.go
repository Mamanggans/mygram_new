package repository

import (
	"context"
	"fmt"
	"mygram-byferdiansyah/domain"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *imageRepository {
	return &imageRepository{db}
}

func (imageRepository *imageRepository) Get(ctx context.Context, images *[]domain.Image) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = imageRepository.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "email")
	}).Find(&images).Error; err != nil {
		return err
	}

	return
}

func (imageRepository *imageRepository) Create(ctx context.Context, image *domain.Image) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	ID, _ := gonanoid.New(16)

	image.ID = fmt.Sprintf("image-%s", ID)

	if err := imageRepository.db.WithContext(ctx).Create(&image).Error; err != nil {
		return err
	}

	return
}

func (imageRepository *imageRepository) GetByID(ctx context.Context, image *domain.Image, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = imageRepository.db.WithContext(ctx).First(&image, &id).Error; err != nil {
		return err
	}

	return
}

func (imageRepository *imageRepository) Edit(ctx context.Context, image domain.Image, id string) (p domain.Image, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	p = domain.Image{}

	if err = imageRepository.db.WithContext(ctx).First(&p, &id).Error; err != nil {
		return p, err
	}

	if err = imageRepository.db.WithContext(ctx).Model(&p).Updates(image).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (imageRepository *imageRepository) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = imageRepository.db.WithContext(ctx).First(&domain.Image{}, &id).Error; err != nil {
		return err
	}

	if err = imageRepository.db.WithContext(ctx).Delete(&domain.Image{}, &id).Error; err != nil {
		return err
	}

	return
}
