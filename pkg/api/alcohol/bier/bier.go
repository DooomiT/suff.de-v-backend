package alcohol

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Bier struct {
	ID string `json:"id"`
	Type  string `json:"type"`
	Brand  string `json:"brand"`
	Description string `json:"description"`
	Image  string `json:"image"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{bierID}", GetABier)
	router.Delete("/{bierID}", DeleteABier)
	router.Post("/", CreateBier)
	router.Get("/", GetAllBiers)
	return router
}

func GetABier(w http.ResponseWriter, r *http.Request) {
	bierID := chi.URLParam(r, "bierID")
	//TODO: implement
	bier := Bier{
		ID: bierID,
		Type:  "Weizen",
		Brand: "Augustiner"
		Description: "Mehr als 30% oder so",
		Image:  "No Image yet ...",
	}
	render.JSON(w, r, bier)
}

func DeleteABier(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	response := make(map[string]string)
	response["message"] = "Deleted Bier successfully"
	render.JSON(w, r, response)
}

func CreateBier(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	response := make(map[string]string)
	response["message"] = "Created Bier successfully"
	render.JSON(w, r, response)
}

func GetAllBiers(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	biers := []Bier{
		{
			ID: bierID,
			Type:  "Weizen",
			Brand: "Augustiner"
			Description: "Mehr als 30% oder so",
			Image:  "No Image yet ...",
		},
		{
			ID: bierID,
			Type:  "Helles",
			Brand: "Schwabenbräu"
			Description: "Mehr als 30% oder so",
			Image:  "No Image yet ...",
		},
	}
	render.JSON(w, r, biers) // A chi router helper for serializing and returning json
}
