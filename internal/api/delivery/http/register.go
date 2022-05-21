package http

import (
	"groupie-tracker-new/internal"
	"net/http"
)

func RegisterHTTPEndpoints(router *http.ServeMux, auc internal.ArtistsUseCase) {
	fs := http.FileServer(http.Dir("./static"))
	h := NewHandler(auc)
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("/", h.Index)
	router.HandleFunc("/artists/", h.GetJson)
	router.HandleFunc("/create/", h.CreateJSON)
}
