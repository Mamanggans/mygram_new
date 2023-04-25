package usecase_test

import (
	"context"
	"errors"
	"mygram-byferdiansyah/domain"
	"mygram-byferdiansyah/domain/mocks"
	"testing"
	"time"

	imageUseCase "mygram-byferdiansyah/image/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	mockImage := domain.Image{
		ID:       "image-123",
		Title:    "A Title",
		Caption:  "A caption",
		ImageUrl: "https://www.example.com/image.jpg",
		UserID:   "user-123",
	}

	mockImages := make([]domain.Image, 0)

	mockImages = append(mockImages, mockImage)

	mockImageRepository := new(mocks.ImageRepository)
	imageUseCase := imageUseCase.NewImageUseCase(mockImageRepository)

	t.Run("get all images correctly", func(t *testing.T) {
		mockImageRepository.On("Get", mock.Anything, mock.AnythingOfType("*[]domain.Image")).Return(nil).Once()

		err := imageUseCase.Get(context.Background(), &mockImages)

		assert.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	now := time.Now()
	mockAddedImage := domain.Image{
		ID:        "image-123",
		Title:     "A Title",
		Caption:   "A caption",
		ImageUrl:  "https://www.example.com/image.jpg",
		UserID:    "user-123",
		CreatedAt: &now,
	}

	mockImageRepository := new(mocks.ImageRepository)
	imageUseCase := imageUseCase.NewImageUseCase(mockImageRepository)

	t.Run("add image correctly", func(t *testing.T) {
		tempMockAddImage := domain.Image{
			Title:    "A Title",
			Caption:  "A caption",
			ImageUrl: "https://www.example.com/image.jpg",
		}

		tempMockAddImage.ID = "image-123"

		mockImageRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Image")).Return(nil).Once()

		err := imageUseCase.Create(context.Background(), &tempMockAddImage)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddImage)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedImage.ID, tempMockAddImage.ID)
		assert.Equal(t, mockAddedImage.Title, tempMockAddImage.Title)
		assert.Equal(t, mockAddedImage.Caption, tempMockAddImage.Caption)
		assert.Equal(t, mockAddedImage.ImageUrl, tempMockAddImage.ImageUrl)
		mockImageRepository.AssertExpectations(t)
	})

	t.Run("add image with empty title", func(t *testing.T) {
		tempMockAddImage := domain.Image{
			Title:    "",
			Caption:  "A caption",
			ImageUrl: "https://www.example.com/image.jpg",
		}

		tempMockAddImage.ID = "image-123"

		mockImageRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Image")).Return(nil).Once()

		err := imageUseCase.Create(context.Background(), &tempMockAddImage)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddImage)

		assert.Error(t, err)
		assert.Equal(t, mockAddedImage.ID, tempMockAddImage.ID)
		assert.NotEqual(t, mockAddedImage.Title, tempMockAddImage.Title)
		assert.Equal(t, mockAddedImage.Caption, tempMockAddImage.Caption)
		assert.Equal(t, mockAddedImage.ImageUrl, tempMockAddImage.ImageUrl)
		mockImageRepository.AssertExpectations(t)
	})

	t.Run("add image with empty image url", func(t *testing.T) {
		tempMockAddImage := domain.Image{
			Title:    "A Title",
			Caption:  "A caption",
			ImageUrl: "",
		}

		tempMockAddImage.ID = "image-123"

		mockImageRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Image")).Return(nil).Once()

		err := imageUseCase.Create(context.Background(), &tempMockAddImage)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddImage)

		assert.Error(t, err)
		assert.Equal(t, mockAddedImage.ID, tempMockAddImage.ID)
		assert.Equal(t, mockAddedImage.Title, tempMockAddImage.Title)
		assert.Equal(t, mockAddedImage.Caption, tempMockAddImage.Caption)
		assert.NotEqual(t, mockAddedImage.ImageUrl, tempMockAddImage.ImageUrl)
		mockImageRepository.AssertExpectations(t)
	})

	t.Run("add image with not contain needed property", func(t *testing.T) {
		tempMockAddImage := domain.Image{
			Title:   "A Title",
			Caption: "A caption",
		}

		tempMockAddImage.ID = "image-123"

		mockImageRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Image")).Return(nil).Once()

		err := imageUseCase.Create(context.Background(), &tempMockAddImage)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddImage)

		assert.Error(t, err)
		assert.Equal(t, mockAddedImage.ID, tempMockAddImage.ID)
		assert.Equal(t, mockAddedImage.Title, tempMockAddImage.Title)
		assert.Equal(t, mockAddedImage.Caption, tempMockAddImage.Caption)
		assert.NotEqual(t, mockAddedImage.ImageUrl, tempMockAddImage.ImageUrl)
		mockImageRepository.AssertExpectations(t)
	})
}

func TestGetBy(t *testing.T) {
	var mockImage *domain.Image

	now := time.Now()

	mockImage = &domain.Image{
		ID:        "image-123",
		Title:     "A Title",
		Caption:   "A caption",
		ImageUrl:  "https://www.example.com/image.jpg",
		UserID:    "user-123",
		CreatedAt: &now,
	}

	mockImageRepository := new(mocks.ImageRepository)
	imageUseCase := imageUseCase.NewImageUseCase(mockImageRepository)

	t.Run("get by id correctly", func(t *testing.T) {
		mockImageID := "image-123"

		mockImageRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.Image"), mock.AnythingOfType("string")).Return(nil).Once()

		err := imageUseCase.GetByID(context.Background(), mockImage, mockImageID)

		assert.NoError(t, err)
		assert.Equal(t, mockImage.ID, mockImageID)
		mockImageRepository.AssertExpectations(t)
	})

	t.Run("get by id with not found image", func(t *testing.T) {
		mockImageID := "image-234"

		mockImageRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.Image"), mock.AnythingOfType("string")).Return(nil).Once()

		err := imageUseCase.GetByID(context.Background(), mockImage, mockImageID)

		assert.NoError(t, err)
		assert.NotEqual(t, mockImage.ID, mockImageID)
		mockImageRepository.AssertExpectations(t)
	})
}

func TestEdit(t *testing.T) {
	now := time.Now()
	mockEditedImage := domain.Image{
		ID:        "image-123",
		Title:     "A New Title",
		Caption:   "A new caption",
		ImageUrl:  "https://www.example.com/new-image.jpg",
		UserID:    "user-123",
		UpdatedAt: &now,
	}

	mockImageRepository := new(mocks.ImageRepository)
	imageUseCase := imageUseCase.NewImageUseCase(mockImageRepository)

	t.Run("edit image correctly", func(t *testing.T) {
		tempMockImageID := "image-123"
		tempMockEditImage := domain.Image{
			Title:    "A New Title",
			Caption:  "A new caption",
			ImageUrl: "https://www.example.com/new-image.jpg",
		}

		mockImageRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.Image"), mock.AnythingOfType("string")).Return(mockEditedImage, nil).Once()

		image, err := imageUseCase.Edit(context.Background(), tempMockEditImage, tempMockImageID)

		assert.NoError(t, err)

		tempMockEditedImage := domain.Image{
			ID:        tempMockImageID,
			Title:     tempMockEditImage.Title,
			Caption:   tempMockEditImage.Caption,
			ImageUrl:  tempMockEditImage.ImageUrl,
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditedImage)

		assert.NoError(t, err)
		assert.Equal(t, image, tempMockEditedImage)
		assert.Equal(t, mockEditedImage.Title, tempMockEditImage.Title)
		assert.Equal(t, mockEditedImage.Caption, tempMockEditImage.Caption)
		assert.Equal(t, mockEditedImage.ImageUrl, tempMockEditedImage.ImageUrl)
		mockImageRepository.AssertExpectations(t)
	})

	t.Run("edit image with empty title", func(t *testing.T) {
		tempMockImageID := "image-123"
		tempMockEditImage := domain.Image{
			Title:    "",
			Caption:  "A new caption",
			ImageUrl: "https://www.example.com/new-image.jpg",
		}

		mockImageRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.Image"), mock.AnythingOfType("string")).Return(mockEditedImage, nil).Once()

		image, err := imageUseCase.Edit(context.Background(), tempMockEditImage, tempMockImageID)

		assert.NoError(t, err)

		tempMockEditedImage := domain.Image{
			ID:        tempMockImageID,
			Title:     tempMockEditImage.Title,
			Caption:   tempMockEditImage.Caption,
			ImageUrl:  tempMockEditImage.ImageUrl,
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditedImage)

		assert.Error(t, err)
		assert.NotEqual(t, image, tempMockEditedImage)
		assert.NotEqual(t, mockEditedImage.Title, tempMockEditImage.Title)
		assert.Equal(t, mockEditedImage.Caption, tempMockEditImage.Caption)
		assert.Equal(t, mockEditedImage.ImageUrl, tempMockEditedImage.ImageUrl)
		mockImageRepository.AssertExpectations(t)
	})

	t.Run("edit image with empty image url", func(t *testing.T) {
		tempMockImageID := "image-123"
		tempMockEditImage := domain.Image{
			Title:    "A New Title",
			Caption:  "A new caption",
			ImageUrl: "",
		}

		mockImageRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.Image"), mock.AnythingOfType("string")).Return(mockEditedImage, nil).Once()

		image, err := imageUseCase.Edit(context.Background(), tempMockEditImage, tempMockImageID)

		assert.NoError(t, err)

		tempMockEditedImage := domain.Image{
			ID:        tempMockImageID,
			Title:     tempMockEditImage.Title,
			Caption:   tempMockEditImage.Caption,
			ImageUrl:  tempMockEditImage.ImageUrl,
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditedImage)

		assert.Error(t, err)
		assert.NotEqual(t, image, tempMockEditedImage)
		assert.Equal(t, mockEditedImage.Title, tempMockEditImage.Title)
		assert.Equal(t, mockEditedImage.Caption, tempMockEditImage.Caption)
		assert.NotEqual(t, mockEditedImage.ImageUrl, tempMockEditedImage.ImageUrl)
		mockImageRepository.AssertExpectations(t)
	})

	t.Run("edit image with empty title and image url", func(t *testing.T) {
		tempMockImageID := "image-123"
		tempMockEditImage := domain.Image{
			Title:    "",
			Caption:  "A new caption",
			ImageUrl: "",
		}

		mockImageRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.Image"), mock.AnythingOfType("string")).Return(mockEditedImage, nil).Once()

		image, err := imageUseCase.Edit(context.Background(), tempMockEditImage, tempMockImageID)

		assert.NoError(t, err)

		tempMockEditedImage := domain.Image{
			ID:        tempMockImageID,
			Title:     tempMockEditImage.Title,
			Caption:   tempMockEditImage.Caption,
			ImageUrl:  tempMockEditImage.ImageUrl,
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditedImage)

		assert.Error(t, err)
		assert.NotEqual(t, image, tempMockEditedImage)
		assert.NotEqual(t, mockEditedImage.Title, tempMockEditImage.Title)
		assert.Equal(t, mockEditedImage.Caption, tempMockEditImage.Caption)
		assert.NotEqual(t, mockEditedImage.ImageUrl, tempMockEditedImage.ImageUrl)
		mockImageRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockImage := domain.Image{
		ID:       "image-123",
		Title:    "A Title",
		Caption:  "A caption",
		ImageUrl: "https://www.example.com/image.jpg",
		UserID:   "user-123",
	}

	mockImageRepository := new(mocks.ImageRepository)
	imageUseCase := imageUseCase.NewImageUseCase(mockImageRepository)

	t.Run("delete image correctly", func(t *testing.T) {
		mockImageRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := imageUseCase.Delete(context.Background(), mockImage.ID)

		assert.NoError(t, err)
		mockImageRepository.AssertExpectations(t)
	})

	t.Run("delete image with not found image", func(t *testing.T) {
		mockImageRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("fail")).Once()

		err := imageUseCase.Delete(context.Background(), "image-234")

		assert.Error(t, err)
		mockImageRepository.AssertExpectations(t)
	})
}
