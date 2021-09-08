package albums

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Album struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{albumID}", GetAAlbum)
	router.Delete("/{albumID}", DeleteAlbum)
	router.Post("/", CreateAlbum)
	router.Get("/", GetAllAlbums)
	return router
}

func GetAAlbum(w http.ResponseWriter, r *http.Request) {
	albumID := chi.URLParam(r, "albumID")
	albums := Album{
		Slug:  albumID,
		Title: "Hello world",
		Body:  "Heloo world from planet earth",
	}
	render.JSON(w, r, albums) // A chi router helper for serializing and returning json
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Album successfully"
	render.JSON(w, r, response) // Return some demo response
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Album successfully"
	render.JSON(w, r, response) // Return some demo response
}

func GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	albums := []Album{
		{
			Slug:  "slug",
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		},
	}
	render.JSON(w, r, albums) // A chi router helper for serializing and returning json
}
