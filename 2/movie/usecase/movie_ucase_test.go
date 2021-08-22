package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	ucase "github.com/bxcodec/go-clean-arch/movie/usecase"
)

func TestGetByID(t *testing.T) {
	var response = "True"

	mockMovieRepo := new(mocks.MovieRepository)
	mockMovie := domain.Movie{
		ImdbID:         "tt4853102",
		Title:          "Batman: The Killing Joke",
		Year:           "2016",
		Type:           "movie",
		Poster:         "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		ResponseStatus: &response,
	}

	t.Run("success", func(t *testing.T) {
		mockMovieRepo.
			On("GetByID", mock.Anything, mock.AnythingOfType("string")).
			Return(mockMovie, nil).
			Once()

		u := ucase.NewMovieUsecase(mockMovieRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockMovie.ImdbID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockMovieRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockMovieRepo.
			On("GetByID", mock.Anything, mock.AnythingOfType("string")).
			Return(domain.Movie{}, errors.New("Unexpected")).Once()

		u := ucase.NewMovieUsecase(mockMovieRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockMovie.ImdbID)

		assert.Error(t, err)
		assert.Equal(t, domain.Movie{}, a)

		mockMovieRepo.AssertExpectations(t)
	})
}

func TestSearch(t *testing.T) {
	var response = "True"

	mockLogRepo := new(mocks.LogRepository)
	mockMovieRepo := new(mocks.MovieRepository)
	mockMovies := make([]domain.Movie, 0)
	mockMovies = append(mockMovies, domain.Movie{
		ImdbID:         "tt4853102",
		Title:          "Batman: The Killing Joke",
		Year:           "2016",
		Type:           "movie",
		Poster:         "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		ResponseStatus: &response,
	})

	dMockMovies := domain.Movies{
		Movie: &mockMovies,
	}

	t.Run("success", func(t *testing.T) {
		mockMovieRepo.
			On("Search", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("string")).
			Return(dMockMovies, nil).
			Once()

		u := ucase.NewMovieUsecase(mockMovieRepo, time.Second*2)

		a, err := u.Search(context.TODO(), 1, "dummy")

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockMovieRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockMovieRepo.
			On("Search", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("string")).
			Return(domain.Movies{}, errors.New("Unexpected")).Once()

		u := ucase.NewMovieUsecase(mockMovieRepo, time.Second*2)

		a, err := u.Search(context.TODO(), 1, "dummy")

		assert.Error(t, err)
		assert.Equal(t, domain.Movies{}, a)

		mockMovieRepo.AssertExpectations(t)
	})
}
