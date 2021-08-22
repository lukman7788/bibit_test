package http

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ResponseError ...
type ResponseError struct {
	Message string `json:"message"`
}

// MovieHandler  ...
type MovieHandler struct {
	AUsecase domain.MovieUsecase
}

// NewMovieHandler ...
func NewHttpMovieHandler(e *echo.Echo, us domain.MovieUsecase) {
	handler := &MovieHandler{
		AUsecase: us,
	}

	e.GET("/movie/:id", handler.GetByID)
	e.GET("/movie-search", handler.Search)
}

// Search ...
func (a *MovieHandler) Search(c echo.Context) error {
	paginationP, err := strconv.Atoi(c.QueryParam("pagination"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrBadParamInput.Error())
	}

	searchword := c.QueryParam("searchword")
	if searchword == "" {
		return c.JSON(http.StatusNotFound, domain.ErrBadParamInput.Error())
	}
	pagination := int64(paginationP)

	ctx := c.Request().Context()

	srch, err := a.AUsecase.Search(ctx, pagination, searchword)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, srch)
}

// GetByID ...
func (a *MovieHandler) GetByID(c echo.Context) error {
	var id string = c.Param("id")
	ctx := c.Request().Context()

	mov, err := a.AUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, mov)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
