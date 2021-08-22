package grpc

import (
	"context"
	"log"

	"github.com/bxcodec/go-clean-arch/domain"
	movie_grpc "github.com/bxcodec/go-clean-arch/movie/delivery/grpc/movie_grpc"
	grpc "google.golang.org/grpc"
)

// GRPCServer  ...
type GRPCServer struct {
	usecase domain.MovieUsecase
}

// NewMovieHandler ...
func NewGrpcMovieHandler(gserver *grpc.Server, us domain.MovieUsecase) {
	handler := &GRPCServer{
		usecase: us,
	}
	movie_grpc.RegisterMovieHandlerServer(gserver, handler)
}

// GetByID ...
func (a *GRPCServer) GetByID(ctx context.Context, req *movie_grpc.MovieRequest) (*movie_grpc.MovieData, error) {
	var ImdbID string = req.ImdbID

	mov, err := a.usecase.GetByID(ctx, ImdbID)
	if err != nil {
		return &movie_grpc.MovieData{}, err
	}

	ratings := make([]*movie_grpc.Rating, 0)
	for _, rating := range *mov.Ratings {
		ratings = append(ratings, &movie_grpc.Rating{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	return &movie_grpc.MovieData{
		ImdbID:   mov.ImdbID,
		Title:    mov.Title,
		Year:     mov.Year,
		Type:     mov.Type,
		Poster:   mov.Poster,
		Rated:    pointerStringHandler(mov.Rated),
		Released: pointerStringHandler(mov.Released),
		Runtime:  pointerStringHandler(mov.Runtime),
		Genre:    pointerStringHandler(mov.Genre),
		Director: pointerStringHandler(mov.Director),
		Writer:   pointerStringHandler(mov.Writer),
		Actors:   pointerStringHandler(mov.Actors),
		Ratings:  ratings,
	}, nil
}

func pointerStringHandler(variable *string) string {
	if &variable != nil {
		return string(*variable)
	}
	return ""
}

// Search ...
func (a *GRPCServer) Search(ctx context.Context, req *movie_grpc.SearchRequest) (*movie_grpc.SearchData, error) {
	var pagination int64 = req.Pagination
	var searchword string = req.Searchword

	search, err := a.usecase.Search(ctx, pagination, searchword)
	if err != nil {
		return &movie_grpc.SearchData{}, err
	}

	searchMovie := make([]*movie_grpc.ShortMovie, 0)
	for _, srch := range *search.Movie {
		log.Println(srch.ImdbID)
		log.Println(srch.Title)
		log.Println(srch.Year)
		log.Println(srch.Type)
		searchMovie = append(searchMovie, &movie_grpc.ShortMovie{
			ImdbID: srch.ImdbID,
			Title:  srch.Title,
			Year:   srch.Year,
			Type:   srch.Type,
		})
	}

	return &movie_grpc.SearchData{
		Search:      searchMovie,
		TotalResult: pointerStringHandler(search.TotalResult),
		Response:    pointerStringHandler(search.ResponseStatus),
	}, nil
}
