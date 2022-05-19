package internal

import (
	"context"
	"groupie-tracker-new/models"
)

type ArtistsUseCase interface {
	GetArtists(ctx context.Context) ([]*models.Artist, error)
	GetArtistById(ctx context.Context, id int) (*models.Artist, error)
	Create(ctx context.Context)
}
