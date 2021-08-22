package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/spf13/viper"
)

type apiMovieRepository struct {
	HTTPClient http.Client
}

// NewAPIMovieRepository ...
func NewAPIMovieRepository() domain.MovieRepository {
	return &apiMovieRepository{}
}

func (m *apiMovieRepository) GetByID(ctx context.Context, ImdbID string) (movie domain.Movie, err error) {
	req, err := http.NewRequest(
		"GET", "http://www.omdbapi.com?apikey="+viper.GetString("imdbAPIKEY")+"&i="+ImdbID, nil,
	)

	if err != nil {
		return domain.Movie{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		return domain.Movie{}, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&movie)

	if *movie.ResponseStatus == "False" {
		return domain.Movie{}, errors.New(*movie.ErrorMessage)
	}

	if err != nil {
		return domain.Movie{}, err
	}

	return movie, err
}

func (m *apiMovieRepository) Search(ctx context.Context, pagination int64, searchword string) (movies domain.Movies, err error) {
	pathURL := "http://www.omdbapi.com/?apikey=" + viper.GetString(`imdbAPIKEY`) + "&s=" + searchword + "&page=" + strconv.FormatInt(pagination, 10)
	req, err := http.NewRequest(
		"GET", pathURL, nil,
	)

	if err != nil {
		return domain.Movies{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := m.HTTPClient.Do(req)
	if err != nil {
		return domain.Movies{}, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&movies)

	if *movies.ResponseStatus == "False" {
		return domain.Movies{}, errors.New(*movies.ErrorMessage)
	}

	if err != nil {
		return domain.Movies{}, err
	}

	return movies, err
}
