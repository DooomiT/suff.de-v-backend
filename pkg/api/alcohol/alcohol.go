package alcohol

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Alcohol struct {
	ID string `json:"id"`
	Type  string `json:"type"`
	Description string `json:"description"`
	Image  string `json:"image"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{alcoholID}", GetAAlcohol)
	router.Delete("/{alcoholID}", DeleteAAlcohol)
	router.Post("/", CreateAlcohol)
	router.Get("/", GetAllAlcohols)
	return router
}

func GetAAlcohol(w http.ResponseWriter, r *http.Request) {
	alcoholID := chi.URLParam(r, "alcoholID")
	//TODO: implement
	alcohol := Alcohol{
		ID: alcoholID,
		Type:  "Hochprozentiges",
		Description: "Mehr als 30% oder so",
		Image:  "No Image yet ...",
	}
	render.JSON(w, r, alcohol)
}

func DeleteAAlcohol(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	response := make(map[string]string)
	response["message"] = "Deleted Album successfully"
	render.JSON(w, r, response)
}

func CreateAlcohol(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	response := make(map[string]string)
	response["message"] = "Created Album successfully"
	render.JSON(w, r, response)
}

func GetAllAlcohols(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	alcohols := []Alcohol{
		{
			ID: "001",
			Type:  "Hochprozentiges",
			Description: "Mehr als 30% oder so",
			Image:  "No Image yet ...",
		},
		{
			ID: "002",
			Type:  "Likör",
			Description: "Weniger als 30% oder so",
			Image:  "No Image yet ...",
		},
	}
	render.JSON(w, r, alcohols) // A chi router helper for serializing and returning json
}
