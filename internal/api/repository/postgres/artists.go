package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"groupie-tracker-new/internal"
	"groupie-tracker-new/models"
	"log"
	"net/http"
)

type ArtistsRepo struct {
	db *sql.DB
}

func NewArtistsRepository(db *sql.DB) internal.ArtistsRepository {
	return &ArtistsRepo{
		db: db,
	}
}

func (ar ArtistsRepo) GetAll(ctx context.Context) ([]*models.Artist, error) {
	// sqlQuery := `SELECT * FROM groupie_tracker_artists`
	// result := []*models.Artist{}
	// rows, err := ar.db.QueryContext(ctx, sqlQuery)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (ar ArtistsRepo) GetOne(ctx context.Context, id int) (*models.Artist, error) {
	return nil, nil
}

func (ar ArtistsRepo) Create(ctx context.Context) {
	fmt.Println("Start create Artists")
	sqlQuery := `INSERT INTO groupie_tracker_artists(artists) VALUES($1)`
	var err error

	artists_information, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Print(err)
	}
	artists := models.Artists{}
	json.NewDecoder(artists_information.Body).Decode(&artists.Artists)
	fmt.Println("Checkpoint1")
	row := ar.db.QueryRowContext(ctx, sqlQuery, artists.Artists)
	err = row.Scan()
	fmt.Println(err)
}
