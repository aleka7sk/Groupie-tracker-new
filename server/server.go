package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"groupie-tracker-new/internal"
	authhttp "groupie-tracker-new/internal/api/delivery/http"
	"groupie-tracker-new/internal/api/repository/postgres"
	"groupie-tracker-new/internal/domain/usecase"
	"groupie-tracker-new/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type App struct {
	httpServer     *http.Server
	Logger         *logrus.Logger
	artistsUseCase internal.ArtistsUseCase
}

func NewApp() *App {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Db initialization error %v", err)
	}
	CreateTables(db)
	artistsRepository := postgres.NewArtistsRepository(db)
	return &App{
		artistsUseCase: usecase.NewService(artistsRepository),
		Logger:         logrus.New(),
	}
}

func CreateTables(db *sql.DB) {
	file, err := os.ReadFile("internal/api/repository/postgres/group.sql")
	fmt.Println(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Exec(string(file))
	if err != nil {
		log.Fatal(err.Error())
	}
	artists_information, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Print(err)
	}
	groups := []models.Group{}
	json.NewDecoder(artists_information.Body).Decode(&groups)
	sqlStatement := `INSERT INTO group_name (ID, IMAGE, NAME, MEMBERS, CREATION_DATE, FIRST_ALBUM, LOCATIONS, CONCERT_DATES, RELATIONS) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	for _, group := range groups {
		_, err = db.Exec(sqlStatement, group.Id, group.Image, group.Name, pq.Array(group.Members), group.CreationDate, group.FirstAlbum, group.Locations, group.ConcertDates, group.Relations)
		if err != nil {
			log.Fatalf("Insert error: %v", err)
		}
	}
}

func (a *App) Run(port string) error {
	router := http.NewServeMux()
	a.Logger.Info("Initialize router...")
	authhttp.RegisterHTTPEndpoints(router, a.artistsUseCase)
	a.Logger.Info("Register HTTP endpoints...")

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			a.Logger.Fatalf("Failed to listen and server: %+v", err)
		}
	}()
	a.Logger.Printf("Server run on port: %v", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.httpServer.Shutdown(ctx)
}

func initDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%s "+
		"password=%s sslmode=disable",
		"localhost", "5432", "postgres", "password")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
