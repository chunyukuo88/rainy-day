package models

type Movie struct {
	ID           int
	TMDB_ID      int
	Title        string
	ChineseTitle *string
	Tagline      string
	ReleaseYear  int
	Genres       []MovieGenre
	Overview     *string
	Score        *float32
	Popularity   *float32
	Keywords     []string
	Language     *string
	PosterURL    *string
	TrailerUrl   *string
	Casting      []Actor
}
