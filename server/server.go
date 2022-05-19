package server

import (
	"context"
	"database/sql"
	"fmt"
	"groupie-tracker-new/internal"
	authhttp "groupie-tracker-new/internal/api/delivery/http"
	"groupie-tracker-new/internal/api/repository/postgres"
	"groupie-tracker-new/internal/domain/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
)

type App struct {
	httpServer     *http.Server
	artistsUseCase internal.ArtistsUseCase
}

func NewApp() *App {
	db, _ := initDB()
	artistsRepository := postgres.NewArtistsRepository(db)
	return &App{
		artistsUseCase: usecase.NewService(artistsRepository),
	}
}

func (a *App) Run(port string) error {
	router := http.NewServeMux()
	authhttp.RegisterHTTPEndpoints(router, a.artistsUseCase)
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and server: %+v", err)
		}
	}()
	fmt.Printf("Server run on port: %v", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.httpServer.Shutdown(ctx)
}

func initDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%s "+
		"password=%s dbname=%s sslmode=require",
		"localhost", "5432", "postgres", "", "groupie_tracker")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
