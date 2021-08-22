package domain

import (
	"context"
)

// Movie ...
type Movie struct {
	ImdbID         string  `json:"imdbID"`
	Title          string  `json:"Title"`
	Year           string  `json:"Year"`
	Type           string  `json:"Type"`
	Poster         string  `json:"Poster"`
	Rated          *string `json:"Rated,omitempty"`
	Released       *string `json:"Released,omitempty"`
	Runtime        *string `json:"Runtime,omitempty"`
	Genre          *string `json:"Genre,omitempty"`
	Director       *string `json:"Director,omitempty"`
	Writer         *string `json:"Writer,omitempty"`
	Actors         *string `json:"Actors,omitempty"`
	Plot           *string `json:"Plot,omitempty"`
	Language       *string `json:"Language,omitempty"`
	Country        *string `json:"Country,omitempty"`
	Awards         *string `json:"Awards,omitempty"`
	Ratings        *[]rate `json:"Ratings,omitempty"`
	Metascore      *string `json:"Metascore,omitempty"`
	ImdbRating     *string `json:"imdbRating,omitempty"`
	ImdbVotes      *string `json:"imdbVotes,omitempty"`
	DVD            *string `json:"DVD,omitempty"`
	BoxOffice      *string `json:"BoxOffice,omitempty"`
	Production     *string `json:"Production,omitempty"`
	Website        *string `json:"Website,omitempty"`
	ResponseStatus *string `json:"Response,omitempty"`
	ErrorMessage   *string `json:"Error,omitempty"`
}

type Movies struct {
	Movie          *[]Movie `json:"Search"`
	TotalResult    *string  `json:"totalResults"`
	ResponseStatus *string  `json:"Response"`
	ErrorMessage   *string  `json:"Error"`
}

type rate struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

// MovieUsecase represent the Movie's usecases
type MovieUsecase interface {
	Search(ctx context.Context, pagination int64, searchword string) (Movies, error)
	GetByID(ctx context.Context, ImdbID string) (Movie, error)
}

// MovieRepository represent the movie's repository contract
type MovieRepository interface {
	Search(ctx context.Context, pagination int64, searchword string) (Movies, error)
	GetByID(ctx context.Context, ImdbID string) (Movie, error)
}
