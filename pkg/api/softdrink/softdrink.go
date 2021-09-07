package softdrink

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Softdrink struct {
	ID string `json:"id"`
	Type  string `json:"type"`
	Description string `json:"description"`
	Image  string `json:"image"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{softdrinkID}", GetASoftdrink)
	router.Delete("/{softdrinkID}", DeleteASoftdrink)
	router.Post("/", CreateSoftdrink)
	router.Get("/", GetAllSoftdrinks)
	return router
}

func GetASoftdrink(w http.ResponseWriter, r *http.Request) {
	softdrinkID := chi.URLParam(r, "softdrinkID")
	//TODO: implement
	softdrink := Softdrink{
		ID: softdrinkID,
		Type:  "Cola",
		Description: "Mehr als 30% oder so",
		Image:  "No Image yet ...",
	}
	render.JSON(w, r, softdrink)
}

func DeleteASoftdrink(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	response := make(map[string]string)
	response["message"] = "Deleted Album successfully"
	render.JSON(w, r, response)
}

func CreateSoftdrink(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	response := make(map[string]string)
	response["message"] = "Created Album successfully"
	render.JSON(w, r, response)
}

func GetAllSoftdrinks(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	softdrinks := []Softdrink{
		{
			ID: "001",
			Type:  "Cola",
			Description: "Mehr als 30% oder so",
			Image:  "No Image yet ...",
		},
		{
			ID: "002",
			Type:  "Fanta",
			Description: "Weniger als 30% oder so",
			Image:  "No Image yet ...",
		},
	}
	render.JSON(w, r, softdrinks) // A chi router helper for serializing and returning json
}
