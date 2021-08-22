package api_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/go-clean-arch/domain"
	apiMovieRepository "github.com/bxcodec/go-clean-arch/movie/repository/api"
)

func TestGetByID(t *testing.T) {
	ImdbID := "tt4853102"
	jsonmock := `{"Title":"Batman: The Killing Joke","Year":"2016","Rated":"R","Released":"25 Jul 2016","Runtime":"76 min","Genre":"Animation, Action, Crime","Director":"Sam Liu","Writer":"Brian Azzarello, Brian Bolland, Bob Kane","Actors":"Kevin Conroy, Mark Hamill, Tara Strong","Plot":"As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.","Language":"English","Country":"United States","Awards":"1 win & 2 nominations","Poster":"https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.4/10"},{"Source":"Rotten Tomatoes","Value":"39%"}],"Metascore":"N/A","imdbRating":"6.4","imdbVotes":"51,677","imdbID":"tt4853102","Type":"movie","DVD":"26 Jul 2016","BoxOffice":"$3,775,000","Production":"Warner Bros.","Website":"N/A","Response":"True"}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.omdbapi.com?apikey="+viper.GetString("imdbAPIKEY")+"&i="+ImdbID,
		httpmock.NewStringResponder(http.StatusOK, jsonmock))

	apiMovieRepo := apiMovieRepository.NewAPIMovieRepository()

	movie, err := apiMovieRepo.GetByID(context.Background(), ImdbID)

	assert.NoError(t, err)

	var expectMovie domain.Movie
	err = json.Unmarshal([]byte(jsonmock), &expectMovie)
	assert.Equal(t, movie, expectMovie)
}

func TestSearch(t *testing.T) {
	searchword := "dummy"
	var pagination int64 = 1
	pathURL := "http://www.omdbapi.com/?apikey=" + viper.GetString(`imdbAPIKEY`) + "&s=" + searchword + "&page=" + strconv.FormatInt(pagination, 10)
	jsonmock := `{"Search":[{"imdbID":"tt0372784","Title":"Batman Begins","Year":"2005","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"},{"imdbID":"tt2975590","Title":"Batman v Superman: Dawn of Justice","Year":"2016","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"},{"imdbID":"tt0096895","Title":"Batman","Year":"1989","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg"},{"imdbID":"tt0103776","Title":"Batman Returns","Year":"1992","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOGZmYzVkMmItM2NiOS00MDI3LWI4ZWQtMTg0YWZkODRkMmViXkEyXkFqcGdeQXVyODY0NzcxNw@@._V1_SX300.jpg"},{"imdbID":"tt0118688","Title":"Batman \u0026 Robin","Year":"1997","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMGQ5YTM1NmMtYmIxYy00N2VmLWJhZTYtN2EwYTY3MWFhOTczXkEyXkFqcGdeQXVyNTA2NTI0MTY@._V1_SX300.jpg"},{"imdbID":"tt0112462","Title":"Batman Forever","Year":"1995","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNDdjYmFiYWEtYzBhZS00YTZkLWFlODgtY2I5MDE0NzZmMDljXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg"},{"imdbID":"tt4116284","Title":"The Lego Batman Movie","Year":"2017","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTcyNTEyOTY0M15BMl5BanBnXkFtZTgwOTAyNzU3MDI@._V1_SX300.jpg"},{"imdbID":"tt0103359","Title":"Batman: The Animated Series","Year":"1992–1995","Type":"series","Poster":"https://m.media-amazon.com/images/M/MV5BOTM3MTRkZjQtYjBkMy00YWE1LTkxOTQtNDQyNGY0YjYzNzAzXkEyXkFqcGdeQXVyOTgwMzk1MTA@._V1_SX300.jpg"},{"imdbID":"tt1569923","Title":"Batman: Under the Red Hood","Year":"2010","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"imdbID":"tt2313197","Title":"Batman: The Dark Knight Returns, Part 1","Year":"2012","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"}],"totalResults":"445","Response":"True","Error":null}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", pathURL,
		httpmock.NewStringResponder(http.StatusOK, jsonmock))

	apiMovieRepo := apiMovieRepository.NewAPIMovieRepository()

	movies, err := apiMovieRepo.Search(context.Background(), pagination, searchword)

	assert.NoError(t, err)

	var expectMovie domain.Movies
	err = json.Unmarshal([]byte(jsonmock), &expectMovie)
	assert.Equal(t, movies, expectMovie)
}
