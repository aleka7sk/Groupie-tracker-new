package http

import (
	"encoding/json"
	"groupie-tracker-new/internal"
	"groupie-tracker-new/internal/domain/usecase"
	"groupie-tracker-new/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var path = "../../templates/"

type Handler struct {
	usecase internal.ArtistsUseCase
}

func NewHandler(usecase internal.ArtistsUseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error404(w)
		return
	}
	tmpl, err := template.ParseFiles(path + "api/layout.html")
	if err != nil {
		log.Print(err)
		Error505(w)
		return
	}

	// result := artists.ParseApi("https://groupietrackers.herokuapp.com/api/artists")
	// fmt.Println(result)
	// artist_id := usecase.ArtistIdInit()
	// artist_id.ParseApi("https://groupietrackers.herokuapp.com/api/relation")

	// data := usecase.ConnectData(artists, artist_id)

	tmpl.Execute(w, nil)
}

func (h *Handler) GetJson(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artists/")
	w.Header().Set("Content-Type", "application/json")
	data := usecase.Data
	var response models.FullInfo
	for _, elem := range data {
		if strconv.Itoa(elem.Artist.Id) == id {
			response = elem
			break
		}
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CreateJSON(w http.ResponseWriter, r *http.Request) {
	h.usecase.Create(r.Context())
}
