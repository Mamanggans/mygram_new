package usecase

import (
	"context"
	"mygram-byferdiansyah/domain"
)

type commentUseCase struct {
	commentRepository domain.CommentRepository
}

func NewCommentUseCase(commentRepository domain.CommentRepository) *commentUseCase {
	return &commentUseCase{commentRepository}
}

func (commentUseCase *commentUseCase) Get(ctx context.Context, comments *[]domain.Comment, userID string) (err error) {
	if err = commentUseCase.commentRepository.Get(ctx, comments, userID); err != nil {
		return err
	}

	return
}

func (commentUseCase *commentUseCase) Create(ctx context.Context, comment *domain.Comment) (err error) {
	if err = commentUseCase.commentRepository.Create(ctx, comment); err != nil {
		return err
	}

	return
}

func (commentUseCase *commentUseCase) GetByID(ctx context.Context, comment *domain.Comment, id string) (err error) {
	if err = commentUseCase.commentRepository.GetByID(ctx, comment, id); err != nil {
		return err
	}

	return
}

func (commentUseCase *commentUseCase) Edit(ctx context.Context, comment domain.Comment, id string) (image domain.Image, err error) {
	if image, err = commentUseCase.commentRepository.Edit(ctx, comment, id); err != nil {
		return image, err
	}

	return image, nil
}

func (commentUseCase *commentUseCase) Delete(ctx context.Context, id string) (err error) {
	if err = commentUseCase.commentRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}
