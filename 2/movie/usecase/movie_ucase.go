package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type movieUsecase struct {
	movieRepo      domain.MovieRepository
	logRepo        domain.LogRepository
	contextTimeout time.Duration
}

// NewMovieUsecase ...
func NewMovieUsecase(a domain.MovieRepository, b domain.LogRepository, timeout time.Duration) domain.MovieUsecase {
	return &movieUsecase{
		movieRepo:      a,
		logRepo:        b,
		contextTimeout: timeout,
	}
}

func (a *movieUsecase) GetByID(c context.Context, ImdbID string) (res domain.Movie, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.movieRepo.GetByID(ctx, ImdbID)
	if err != nil {
		return
	}

	return res, err
}

func (a *movieUsecase) Search(c context.Context, pagination int64, searchword string) (res domain.Movies, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.movieRepo.Search(ctx, pagination, searchword)
	if err != nil {
		return
	}

	go func() {
		jsonString, err := json.Marshal(res)
		if err == nil {
			err = a.logRepo.Store(context.Background(), domain.Log{
				Pagination: pagination,
				SearchWord: searchword,
				Response:   string(jsonString),
			})
		}
	}()

	return res, err
}
