package http_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
	movieHttp "github.com/bxcodec/go-clean-arch/movie/delivery/http"
)

func TestGetByID(t *testing.T) {
	var mockMovie domain.Movie
	err := faker.FakeData(&mockMovie)
	assert.NoError(t, err)

	mockUCase := new(mocks.MovieUsecase)

	ID := mockMovie.ImdbID

	mockUCase.On("GetByID", mock.Anything, ID).Return(mockMovie, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/movie/"+ID, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("movie/:id")
	c.SetParamNames("id")
	c.SetParamValues(ID)
	handler := movieHttp.MovieHandler{
		AUsecase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestSearch(t *testing.T) {
	var mockMovies domain.Movies
	err := faker.FakeData(&mockMovies)
	assert.NoError(t, err)

	mockUCase := new(mocks.MovieUsecase)

	mockUCase.On("Search", mock.Anything, mock.Anything, mock.Anything).Return(mockMovies, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/movie-search?pagination=1&searchword=spiderman", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("movie-search")
	handler := movieHttp.MovieHandler{
		AUsecase: mockUCase,
	}
	err = handler.Search(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
