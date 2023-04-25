package usecase

import (
	"context"
	"mygram-byferdiansyah/domain"
)

type socialMediaUseCase struct {
	socialMediaRepository domain.SocialMediaRepository
}

func NewSocialMediaUseCase(socialMediaRepository domain.SocialMediaRepository) *socialMediaUseCase {
	return &socialMediaUseCase{socialMediaRepository}
}

func (socialMediaUseCase *socialMediaUseCase) Get(ctx context.Context, socialMedias *[]domain.SocialMedia, userID string) (err error) {
	if err = socialMediaUseCase.socialMediaRepository.Get(ctx, socialMedias, userID); err != nil {
		return err
	}

	return
}

func (socialMediaUseCase *socialMediaUseCase) Create(ctx context.Context, socialMedia *domain.SocialMedia) (err error) {
	if err = socialMediaUseCase.socialMediaRepository.Create(ctx, socialMedia); err != nil {
		return err
	}

	return
}

func (socialMediaUseCase *socialMediaUseCase) GetByID(ctx context.Context, socialMedia *domain.SocialMedia, id string) (err error) {
	if err = socialMediaUseCase.socialMediaRepository.GetByID(ctx, socialMedia, id); err != nil {
		return err
	}

	return
}

func (socialMediaUseCase *socialMediaUseCase) Edit(ctx context.Context, socialMedia domain.SocialMedia, id string) (socmed domain.SocialMedia, err error) {
	if socmed, err = socialMediaUseCase.socialMediaRepository.Edit(ctx, socialMedia, id); err != nil {
		return socmed, err
	}

	return socmed, nil
}

func (socialMediaUseCase *socialMediaUseCase) Delete(ctx context.Context, id string) (err error) {
	if err = socialMediaUseCase.socialMediaRepository.Delete(ctx, id); err != nil {
		return err
	}

	return
}
