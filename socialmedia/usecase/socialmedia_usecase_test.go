package usecase_test

import (
	"context"
	"errors"
	"mygram-byferdiansyah/domain"
	"mygram-byferdiansyah/domain/mocks"
	"testing"
	"time"

	socialMediaUseCase "mygram-byferdiansyah/socialmedia/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	now := time.Now()
	mockSocialMedia := domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "Example",
		SocialMediaUrl: "https://www.example.com/johndoe",
		UserID:         "user-123",
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}

	mockSocialMedias := make([]domain.SocialMedia, 0)

	mockSocialMedias = append(mockSocialMedias, mockSocialMedia)

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("get all social media correctly", func(t *testing.T) {
		mockSocialMediaRepository.On("Get", mock.Anything, mock.AnythingOfType("*[]domain.SocialMedia"), mock.AnythingOfType("string")).Return(nil).Once()

		err := socialMediaUseCase.Get(context.Background(), &mockSocialMedias, mockSocialMedia.UserID)

		assert.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	now := time.Now()
	mockAddedSocialMedia := domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "Example",
		SocialMediaUrl: "https://www.example.com/johndoe",
		UserID:         "user-123",
		CreatedAt:      &now,
	}

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("add social media correctly", func(t *testing.T) {
		tempMockAddSocialMedia := domain.SocialMedia{
			Name:           "Example",
			SocialMediaUrl: "https://www.example.com/johndoe",
		}

		tempMockAddSocialMedia.ID = "socialmedia-123"

		mockSocialMediaRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.SocialMedia")).Return(nil).Once()

		err := socialMediaUseCase.Create(context.Background(), &tempMockAddSocialMedia)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddSocialMedia)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedSocialMedia.ID, tempMockAddSocialMedia.ID)
		assert.Equal(t, mockAddedSocialMedia.Name, tempMockAddSocialMedia.Name)
		assert.Equal(t, mockAddedSocialMedia.SocialMediaUrl, tempMockAddSocialMedia.SocialMediaUrl)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("add social media with empty name", func(t *testing.T) {
		tempMockAddSocialMedia := domain.SocialMedia{
			Name:           "",
			SocialMediaUrl: "https://www.example.com/johndoe",
		}

		tempMockAddSocialMedia.ID = "socialmedia-123"

		mockSocialMediaRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.SocialMedia")).Return(nil).Once()

		err := socialMediaUseCase.Create(context.Background(), &tempMockAddSocialMedia)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddSocialMedia)

		assert.Error(t, err)
		assert.Equal(t, mockAddedSocialMedia.ID, tempMockAddSocialMedia.ID)
		assert.NotEqual(t, mockAddedSocialMedia.Name, tempMockAddSocialMedia.Name)
		assert.Equal(t, mockAddedSocialMedia.SocialMediaUrl, tempMockAddSocialMedia.SocialMediaUrl)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("add social media with empty social media url", func(t *testing.T) {
		tempMockAddSocialMedia := domain.SocialMedia{
			Name:           "Example",
			SocialMediaUrl: "",
		}

		tempMockAddSocialMedia.ID = "socialmedia-123"

		mockSocialMediaRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.SocialMedia")).Return(errors.New("fail")).Once()

		err := socialMediaUseCase.Create(context.Background(), &tempMockAddSocialMedia)

		assert.Error(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddSocialMedia)

		assert.Error(t, err)
		assert.Equal(t, mockAddedSocialMedia.ID, tempMockAddSocialMedia.ID)
		assert.Equal(t, mockAddedSocialMedia.Name, tempMockAddSocialMedia.Name)
		assert.NotEqual(t, mockAddedSocialMedia.SocialMediaUrl, tempMockAddSocialMedia.SocialMediaUrl)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("add social media with not contain needed property", func(t *testing.T) {
		tempMockAddSocialMedia := domain.SocialMedia{
			Name: "Example",
		}

		tempMockAddSocialMedia.ID = "socialmedia-123"

		mockSocialMediaRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.SocialMedia")).Return(nil).Once()

		err := socialMediaUseCase.Create(context.Background(), &tempMockAddSocialMedia)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddSocialMedia)

		assert.Error(t, err)
		assert.Equal(t, mockAddedSocialMedia.ID, tempMockAddSocialMedia.ID)
		assert.Equal(t, mockAddedSocialMedia.Name, tempMockAddSocialMedia.Name)
		assert.NotEqual(t, mockAddedSocialMedia.SocialMediaUrl, tempMockAddSocialMedia.SocialMediaUrl)
		mockSocialMediaRepository.AssertExpectations(t)
	})
}

func TestGetBy(t *testing.T) {
	var mockSocialMedia *domain.SocialMedia

	now := time.Now()

	mockSocialMedia = &domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "Example",
		SocialMediaUrl: "https://www.example.com/johndoe",
		UserID:         "user-123",
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("get by id correctly", func(t *testing.T) {
		mockSocialMediaID := "socialmedia-123"

		mockSocialMediaRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.SocialMedia"), mock.AnythingOfType("string")).Return(nil).Once()

		err := socialMediaUseCase.GetByID(context.Background(), mockSocialMedia, mockSocialMediaID)

		assert.NoError(t, err)
		assert.Equal(t, mockSocialMedia.ID, mockSocialMediaID)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("get by id with not found social media", func(t *testing.T) {
		mockSocialMediaID := "socialmedia-234"

		mockSocialMediaRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.SocialMedia"), mock.AnythingOfType("string")).Return(nil).Once()

		err := socialMediaUseCase.GetByID(context.Background(), mockSocialMedia, mockSocialMediaID)

		assert.NoError(t, err)
		assert.NotEqual(t, mockSocialMedia.ID, mockSocialMediaID)
		mockSocialMediaRepository.AssertExpectations(t)
	})
}

func TestEdit(t *testing.T) {
	now := time.Now()
	mockEditedSocialMedia := domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "New Example",
		SocialMediaUrl: "https://www.newexample.com/johndoe",
		UserID:         "user-123",
		UpdatedAt:      &now,
	}

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("edit social media correctly", func(t *testing.T) {
		tempMockSocialMediaID := "socialmedia-123"
		tempMockEditSocialMedia := domain.SocialMedia{
			Name:           "New Example",
			SocialMediaUrl: "https://www.newexample.com/johndoe",
		}

		mockSocialMediaRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.SocialMedia"), mock.AnythingOfType("string")).Return(mockEditedSocialMedia, nil).Once()

		socialmedia, err := socialMediaUseCase.Edit(context.Background(), tempMockEditSocialMedia, tempMockSocialMediaID)

		assert.NoError(t, err)

		tempMockEditedSocialMedia := domain.SocialMedia{
			ID:             tempMockSocialMediaID,
			Name:           tempMockEditSocialMedia.Name,
			SocialMediaUrl: tempMockEditSocialMedia.SocialMediaUrl,
			UserID:         "user-123",
			UpdatedAt:      &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditSocialMedia)

		assert.NoError(t, err)
		assert.Equal(t, socialmedia, tempMockEditedSocialMedia)
		assert.Equal(t, mockEditedSocialMedia.ID, tempMockEditedSocialMedia.ID)
		assert.Equal(t, mockEditedSocialMedia.Name, tempMockEditedSocialMedia.Name)
		assert.Equal(t, mockEditedSocialMedia.SocialMediaUrl, tempMockEditedSocialMedia.SocialMediaUrl)
		assert.Equal(t, mockEditedSocialMedia.UserID, tempMockEditedSocialMedia.UserID)
		assert.Equal(t, mockEditedSocialMedia.UpdatedAt, tempMockEditedSocialMedia.UpdatedAt)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("edit social media with empty name", func(t *testing.T) {
		tempMockSocialMediaID := "socialmedia-123"
		tempMockEditSocialMedia := domain.SocialMedia{
			Name:           "",
			SocialMediaUrl: "https://www.newexample.com/johndoe",
		}

		mockSocialMediaRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.SocialMedia"), mock.AnythingOfType("string")).Return(mockEditedSocialMedia, nil).Once()

		socialmedia, err := socialMediaUseCase.Edit(context.Background(), tempMockEditSocialMedia, tempMockSocialMediaID)

		assert.NoError(t, err)

		tempMockEditedSocialMedia := domain.SocialMedia{
			ID:             tempMockSocialMediaID,
			Name:           tempMockEditSocialMedia.Name,
			SocialMediaUrl: tempMockEditSocialMedia.SocialMediaUrl,
			UserID:         "user-123",
			UpdatedAt:      &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditSocialMedia)

		assert.Error(t, err)
		assert.NotEqual(t, socialmedia, tempMockEditedSocialMedia)
		assert.Equal(t, mockEditedSocialMedia.ID, tempMockEditedSocialMedia.ID)
		assert.NotEqual(t, mockEditedSocialMedia.Name, tempMockEditedSocialMedia.Name)
		assert.Equal(t, mockEditedSocialMedia.SocialMediaUrl, tempMockEditedSocialMedia.SocialMediaUrl)
		assert.Equal(t, mockEditedSocialMedia.UserID, tempMockEditedSocialMedia.UserID)
		assert.Equal(t, mockEditedSocialMedia.UpdatedAt, tempMockEditedSocialMedia.UpdatedAt)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("edit social media with empty social media url", func(t *testing.T) {
		tempMockSocialMediaID := "socialmedia-123"
		tempMockEditSocialMedia := domain.SocialMedia{
			Name:           "New Example",
			SocialMediaUrl: "",
		}

		mockSocialMediaRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.SocialMedia"), mock.AnythingOfType("string")).Return(mockEditedSocialMedia, nil).Once()

		socialmedia, err := socialMediaUseCase.Edit(context.Background(), tempMockEditSocialMedia, tempMockSocialMediaID)

		assert.NoError(t, err)

		tempMockEditedSocialMedia := domain.SocialMedia{
			ID:             tempMockSocialMediaID,
			Name:           tempMockEditSocialMedia.Name,
			SocialMediaUrl: tempMockEditSocialMedia.Name,
			UserID:         "user-123",
			UpdatedAt:      &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditSocialMedia)

		assert.Error(t, err)
		assert.NotEqual(t, socialmedia, tempMockEditedSocialMedia)
		assert.Equal(t, mockEditedSocialMedia.ID, tempMockEditedSocialMedia.ID)
		assert.Equal(t, mockEditedSocialMedia.Name, tempMockEditedSocialMedia.Name)
		assert.NotEqual(t, mockEditedSocialMedia.SocialMediaUrl, tempMockEditedSocialMedia.SocialMediaUrl)
		assert.Equal(t, mockEditedSocialMedia.UserID, tempMockEditedSocialMedia.UserID)
		assert.Equal(t, mockEditedSocialMedia.UpdatedAt, tempMockEditedSocialMedia.UpdatedAt)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("edit social media with not contain property", func(t *testing.T) {
		tempMockSocialMediaID := "socialmedia-123"
		tempMockEditSocialMedia := domain.SocialMedia{
			Name: "New Example",
		}

		mockSocialMediaRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.SocialMedia"), mock.AnythingOfType("string")).Return(mockEditedSocialMedia, nil).Once()

		socialmedia, err := socialMediaUseCase.Edit(context.Background(), tempMockEditSocialMedia, tempMockSocialMediaID)

		assert.NoError(t, err)

		tempMockEditedSocialMedia := domain.SocialMedia{
			ID:             tempMockSocialMediaID,
			Name:           tempMockEditSocialMedia.Name,
			SocialMediaUrl: tempMockEditSocialMedia.Name,
			UserID:         "user-123",
			UpdatedAt:      &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditSocialMedia)

		assert.Error(t, err)
		assert.NotEqual(t, socialmedia, tempMockEditedSocialMedia)
		assert.Equal(t, mockEditedSocialMedia.ID, tempMockEditedSocialMedia.ID)
		assert.Equal(t, mockEditedSocialMedia.Name, tempMockEditedSocialMedia.Name)
		assert.NotEqual(t, mockEditedSocialMedia.SocialMediaUrl, tempMockEditedSocialMedia.SocialMediaUrl)
		assert.Equal(t, mockEditedSocialMedia.UserID, tempMockEditedSocialMedia.UserID)
		assert.Equal(t, mockEditedSocialMedia.UpdatedAt, tempMockEditedSocialMedia.UpdatedAt)
		mockSocialMediaRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	now := time.Now()
	mockSocialMedia := domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "Example",
		SocialMediaUrl: "https://www.example.com/johndoe",
		UserID:         "user-123",
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("delete social media correctly", func(t *testing.T) {
		mockSocialMediaRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := socialMediaUseCase.Delete(context.Background(), mockSocialMedia.ID)

		assert.NoError(t, err)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("delete social media with not found social media", func(t *testing.T) {
		mockSocialMediaRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("fail")).Once()

		err := socialMediaUseCase.Delete(context.Background(), "socialmedia-234")

		assert.Error(t, err)
		mockSocialMediaRepository.AssertExpectations(t)
	})
}
