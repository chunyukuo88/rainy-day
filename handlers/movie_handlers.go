package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chunyukuo88/rainy-day/logger"
	"github.com/chunyukuo88/rainy-day/models"
)

type MovieHandler struct {
	logger logger.Logger
}

func (mh *MovieHandler) writeJSONResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		mh.logger.Error("Error: Failed to encode movies: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (mh *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
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
	mh.writeJSONResponse(w, movies)
}
