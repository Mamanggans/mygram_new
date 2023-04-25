package usecase_test

import (
	"context"
	"errors"
	"mygram-byferdiansyah/domain"
	"mygram-byferdiansyah/domain/mocks"
	"testing"
	"time"

	commentUseCase "mygram-byferdiansyah/comment/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	now := time.Now()
	mockComment := domain.Comment{
		ID:        "comment-123",
		UserID:    "user-123",
		ImageID:   "image-123",
		Message:   "A message",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	mockComments := make([]domain.Comment, 0)

	mockComments = append(mockComments, mockComment)

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("get all comments correctly", func(t *testing.T) {
		mockCommentRepository.On("Get", mock.Anything, mock.AnythingOfType("*[]domain.Comment"), mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.Get(context.Background(), &mockComments, mockComment.UserID)

		assert.NoError(t, err)
	})
}

func TestCreate(t *testing.T) {
	now := time.Now()
	mockAddedComment := domain.Comment{
		ID:        "comment-123",
		UserID:    "user-123",
		ImageID:   "image-123",
		Message:   "A comment",
		CreatedAt: &now,
	}

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("add comment correctly", func(t *testing.T) {
		tempMockAddComment := domain.Comment{
			Message: "test comment",
			ImageID: "testimg-123",
		}

		tempMockAddComment.ID = "comment-123"

		mockCommentRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Comment")).Return(nil).Once()

		err := commentUseCase.Create(context.Background(), &tempMockAddComment)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddComment)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedComment.ID, tempMockAddComment.ID)
		assert.Equal(t, mockAddedComment.Message, tempMockAddComment.Message)
		assert.Equal(t, mockAddedComment.ImageID, tempMockAddComment.ImageID)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("add comment with empty message", func(t *testing.T) {
		tempMockAddComment := domain.Comment{
			Message: "",
			ImageID: "image-123",
		}

		tempMockAddComment.ID = "comment-123"

		mockCommentRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Comment")).Return(nil).Once()

		err := commentUseCase.Create(context.Background(), &tempMockAddComment)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddComment)

		assert.Error(t, err)
		assert.Equal(t, mockAddedComment.ID, tempMockAddComment.ID)
		assert.NotEqual(t, mockAddedComment.Message, tempMockAddComment.Message)
		assert.Equal(t, mockAddedComment.ImageID, tempMockAddComment.ImageID)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("add comment with empty image id", func(t *testing.T) {
		tempMockAddComment := domain.Comment{
			Message: "A comment",
			ImageID: "",
		}

		tempMockAddComment.ID = "comment-123"

		mockCommentRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Comment")).Return(errors.New("fail")).Once()

		err := commentUseCase.Create(context.Background(), &tempMockAddComment)

		assert.Error(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddComment)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedComment.ID, tempMockAddComment.ID)
		assert.Equal(t, mockAddedComment.Message, tempMockAddComment.Message)
		assert.NotEqual(t, mockAddedComment.ImageID, tempMockAddComment.ImageID)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("add comment with not contain needed property", func(t *testing.T) {
		tempMockAddComment := domain.Comment{
			ImageID: "image-123",
		}

		tempMockAddComment.ID = "comment-123"

		mockCommentRepository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Comment")).Return(nil).Once()

		err := commentUseCase.Create(context.Background(), &tempMockAddComment)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddComment)

		assert.Error(t, err)
		assert.Equal(t, mockAddedComment.ID, tempMockAddComment.ID)
		assert.NotEqual(t, mockAddedComment.Message, tempMockAddComment.Message)
		assert.Equal(t, mockAddedComment.ImageID, tempMockAddComment.ImageID)
		mockCommentRepository.AssertExpectations(t)
	})
}

func TestGetBy(t *testing.T) {
	var mockComment *domain.Comment

	now := time.Now()

	mockComment = &domain.Comment{
		ID:        "comment-123",
		UserID:    "user-123",
		ImageID:   "image-123",
		Message:   "A message",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("get by id correctly", func(t *testing.T) {
		mockCommentID := "comment-123"

		mockCommentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.Comment"), mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.GetByID(context.Background(), mockComment, mockCommentID)

		assert.NoError(t, err)
		assert.Equal(t, mockComment.ID, mockCommentID)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("get by id with not found comment", func(t *testing.T) {
		mockCommentID := "comment-234"

		mockCommentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.Comment"), mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.GetByID(context.Background(), mockComment, mockCommentID)

		assert.NoError(t, err)
		assert.NotEqual(t, mockComment.ID, mockCommentID)
		mockCommentRepository.AssertExpectations(t)
	})
}

func TestEdit(t *testing.T) {
	now := time.Now()
	mockEditedComment := domain.Image{
		ID:        "image-123",
		Title:     "A Title",
		Caption:   "A caption",
		ImageUrl:  "https://www.example.com/image.jpg",
		UserID:    "user-123",
		UpdatedAt: &now,
	}

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("edit comment correctly", func(t *testing.T) {
		tempMockCommentID := "comment-123"
		tempMockEditComment := domain.Comment{
			Message: "A new comment",
		}

		mockCommentRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.Comment"), mock.AnythingOfType("string")).Return(mockEditedComment, nil).Once()

		comment, err := commentUseCase.Edit(context.Background(), tempMockEditComment, tempMockCommentID)

		assert.NoError(t, err)

		tempMockEditedComment := domain.Image{
			ID:        "image-123",
			Title:     "A Title",
			Caption:   "A caption",
			ImageUrl:  "https://www.example.com/image.jpg",
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditComment)

		assert.NoError(t, err)
		assert.Equal(t, comment, tempMockEditedComment)
		assert.Equal(t, mockEditedComment.ID, tempMockEditedComment.ID)
		assert.Equal(t, mockEditedComment.Title, tempMockEditedComment.Title)
		assert.Equal(t, mockEditedComment.Caption, tempMockEditedComment.Caption)
		assert.Equal(t, mockEditedComment.ImageUrl, tempMockEditedComment.ImageUrl)
		assert.Equal(t, mockEditedComment.UserID, tempMockEditedComment.UserID)
		assert.Equal(t, mockEditedComment.UpdatedAt, tempMockEditedComment.UpdatedAt)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("edit comment with empty message", func(t *testing.T) {
		tempMockCommentID := "Comment-123"
		tempMockEditComment := domain.Comment{
			Message: "",
		}

		mockCommentRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.Comment"), mock.AnythingOfType("string")).Return(mockEditedComment, nil).Once()

		comment, err := commentUseCase.Edit(context.Background(), tempMockEditComment, tempMockCommentID)

		assert.NoError(t, err)

		tempMockEditedComment := domain.Image{
			ID:        "image-123",
			Title:     "A Title",
			Caption:   "A caption",
			ImageUrl:  "https://www.example.com/image.jpg",
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditComment)

		assert.Error(t, err)
		assert.Equal(t, comment, tempMockEditedComment)
		assert.Equal(t, mockEditedComment.ID, tempMockEditedComment.ID)
		assert.Equal(t, mockEditedComment.Title, tempMockEditedComment.Title)
		assert.Equal(t, mockEditedComment.Caption, tempMockEditedComment.Caption)
		assert.Equal(t, mockEditedComment.ImageUrl, tempMockEditedComment.ImageUrl)
		assert.Equal(t, mockEditedComment.UserID, tempMockEditedComment.UserID)
		assert.Equal(t, mockEditedComment.UpdatedAt, tempMockEditedComment.UpdatedAt)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("edit comment with not contain property", func(t *testing.T) {
		tempMockCommentID := "comment-123"
		tempMockEditComment := domain.Comment{}

		mockCommentRepository.On("Edit", mock.Anything, mock.AnythingOfType("domain.Comment"), mock.AnythingOfType("string")).Return(mockEditedComment, nil).Once()

		comment, err := commentUseCase.Edit(context.Background(), tempMockEditComment, tempMockCommentID)

		assert.NoError(t, err)

		tempMockEditedComment := domain.Image{
			ID:        "image-123",
			Title:     "A Title",
			Caption:   "A caption",
			ImageUrl:  "https://www.example.com/image.jpg",
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockEditComment)

		assert.Error(t, err)
		assert.Equal(t, comment, tempMockEditedComment)
		assert.Equal(t, mockEditedComment.ID, tempMockEditedComment.ID)
		assert.Equal(t, mockEditedComment.Title, tempMockEditedComment.Title)
		assert.Equal(t, mockEditedComment.Caption, tempMockEditedComment.Caption)
		assert.Equal(t, mockEditedComment.ImageUrl, tempMockEditedComment.ImageUrl)
		assert.Equal(t, mockEditedComment.UserID, tempMockEditedComment.UserID)
		assert.Equal(t, mockEditedComment.UpdatedAt, tempMockEditedComment.UpdatedAt)
		mockCommentRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	now := time.Now()
	mockComment := domain.Comment{
		ID:        "comment-123",
		UserID:    "user-123",
		ImageID:   "image-123",
		Message:   "A message",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("delete comment correctly", func(t *testing.T) {
		mockCommentRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.Delete(context.Background(), mockComment.ID)

		assert.NoError(t, err)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("delete comment with not found Comment", func(t *testing.T) {
		mockCommentRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("fail")).Once()

		err := commentUseCase.Delete(context.Background(), "comment-234")

		assert.Error(t, err)
		mockCommentRepository.AssertExpectations(t)
	})
}
