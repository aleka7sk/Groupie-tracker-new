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

func (h *service) GetArtists(ctx context.Context) ([]*models.Artist, error) {
	return h.repository.GetAll(ctx)
}

func (h *service) GetArtistById(ctx context.Context, id int) (*models.Artist, error) {
	return h.repository.GetOne(ctx, 1)
}

func (h *service) Create(ctx context.Context) {
	h.repository.Create(ctx)
}
