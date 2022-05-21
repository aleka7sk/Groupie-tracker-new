package internal

import (
	"context"
	"groupie-tracker-new/models"
)

type ArtistsRepository interface {
	GetAll(ctx context.Context) (*models.Groups, error)
	GetOne(ctx context.Context, id int) (*models.Group, error)
	Create(ctx context.Context)
}
