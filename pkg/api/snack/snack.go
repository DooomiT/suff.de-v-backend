package snack

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Snack struct {
	ID string `json:"id"`
	Type  string `json:"type"`
	Description string `json:"description"`
	Image  string `json:"image"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{snackID}", GetASnack)
	router.Delete("/{snackID}", DeleteASnack)
	router.Post("/", CreateSnack)
	router.Get("/", GetAllSnacks)
	return router
}

func GetASnack(w http.ResponseWriter, r *http.Request) {
	snackID := chi.URLParam(r, "snackID")
	//TODO: implement
	snack := Snack{
		ID: snackID
		Type:  "Choclate",
		Description: "Mehr als 30% oder so",
		Image:  "No Image yet ...",
	}
	render.JSON(w, r, snack)
}

func DeleteASnack(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	response := make(map[string]string)
	response["message"] = "Deleted Album successfully"
	render.JSON(w, r, response)
}

func CreateSnack(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	response := make(map[string]string)
	response["message"] = "Created Album successfully"
	render.JSON(w, r, response)
}

func GetAllSnacks(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	snacks := []Snack{
		{
			ID: snackID
			Type:  "Choclate",
			Description: "Mehr als 30% oder so",
			Image:  "No Image yet ...",
		},
		{
			ID: snackID
			Type:  "Cookie",
			Description: "Weniger als 30% oder so",
			Image:  "No Image yet ...",
		},
	}
	render.JSON(w, r, snacks) // A chi router helper for serializing and returning json
}
