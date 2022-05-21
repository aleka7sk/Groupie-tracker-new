package postgres

import (
	"context"
	"database/sql"
	"groupie-tracker-new/internal"
	"groupie-tracker-new/models"
	"log"

	"github.com/lib/pq"
)

type ArtistsRepo struct {
	db *sql.DB
}

func NewArtistsRepository(db *sql.DB) internal.ArtistsRepository {
	return &ArtistsRepo{
		db: db,
	}
}

func (ar ArtistsRepo) GetAll(ctx context.Context) (*models.Groups, error) {
	groups := models.Groups{}
	sqlQuery := `SELECT * FROM group_name`
	rows, err := ar.db.Query(sqlQuery)
	if err != nil {
		log.Fatalf("Select query %v", err)
	}
	for rows.Next() {
		group := models.Group{}
		if err := rows.Scan(&group.Id, &group.Image, &group.Name, pq.Array(&group.Members), &group.CreationDate, &group.FirstAlbum, &group.Locations, &group.ConcertDates, &group.Relations); err != nil {
			panic(err)
		}
		groups.Groups = append(groups.Groups, group)

	}

	return &groups, nil
}

func (ar ArtistsRepo) GetOne(ctx context.Context, id int) (*models.Group, error) {
	return nil, nil
}

func (ar ArtistsRepo) Create(ctx context.Context) {
}
