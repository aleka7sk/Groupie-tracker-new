package internal

import (
	"context"
	"groupie-tracker-new/models"
)

type ArtistsRepository interface {
	GetAll(ctx context.Context) ([]*models.Artist, error)
	GetOne(ctx context.Context, id int) (*models.Artist, error)
	Create(ctx context.Context)
}
