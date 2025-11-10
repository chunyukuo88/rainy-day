package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chunyukuo88/rainy-day/models"
)

type MovieHandler struct {
}

func (mh MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     1000,
			Title:       "Once Upon a Time in China",
			Tagline:     "He fought. A lot.",
			ReleaseYear: 1992,
			Genres: []models.MovieGenre{
				{ID: 1, Name: "Action"},
			},
			Keywords: []string{"Jet Li", "Kung Fu"},
			Casting: []models.Actor{
				{
					ID:        1,
					FirstName: "Jet",
					LastName:  "Li",
				},
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
