package usecase

import (
	"context"
	"mygram-byferdiansyah/domain"
)

type imageUseCase struct {
	imageRepository domain.ImageRepository
}

func NewImageUseCase(imageRepository domain.ImageRepository) *imageUseCase {
	return &imageUseCase{imageRepository}
}

func (imageUseCase *imageUseCase) Get(ctx context.Context, images *[]domain.Image) (err error) {
	if err = imageUseCase.imageRepository.Get(ctx, images); err != nil {
		return err
	}

	return
}

func (imageUseCase *imageUseCase) Create(ctx context.Context, image *domain.Image) (err error) {
	if err = imageUseCase.imageRepository.Create(ctx, image); err != nil {
		return err
	}

	return
}

func (imageUseCase *imageUseCase) GetByID(ctx context.Context, image *domain.Image, id string) (err error) {
	if err = imageUseCase.imageRepository.GetByID(ctx, image, id); err != nil {
		return err
	}

	return
}

func (imageUseCase *imageUseCase) Edit(ctx context.Context, image domain.Image, id string) (p domain.Image, err error) {
	if p, err = imageUseCase.imageRepository.Edit(ctx, image, id); err != nil {
		return p, err
	}

	return p, nil
}

func (imageUseCase *imageUseCase) Delete(ctx context.Context, id string) (err error) {
	if err = imageUseCase.imageRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}
