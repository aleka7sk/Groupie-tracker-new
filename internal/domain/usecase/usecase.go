package usecase

import (
	"context"
	"groupie-tracker-new/internal"
	"groupie-tracker-new/models"
)

var Data = []models.FullInfo{}

type service struct {
	repository internal.ArtistsRepository
}

func NewService(repository internal.ArtistsRepository) internal.ArtistsUseCase {
	return &service{repository: repository}
}

func (h *service) GetGroups(ctx context.Context) (*models.Groups, error) {
	return h.repository.GetAll(ctx)
}

func (h *service) GetGroupById(ctx context.Context, id int) (*models.Group, error) {
	return h.repository.GetOne(ctx, 1)
}

func (h *service) Create(ctx context.Context) {
	h.repository.Create(ctx)
}
