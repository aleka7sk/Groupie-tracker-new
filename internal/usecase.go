package internal

import (
	"context"
	"groupie-tracker-new/models"
)

type ArtistsUseCase interface {
	GetGroups(ctx context.Context) (*models.Groups, error)
	GetGroupById(ctx context.Context, id int) (*models.Group, error)
	Create(ctx context.Context)
}
