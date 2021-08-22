package mysql_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/bxcodec/go-clean-arch/domain"

	mysqlLogRepository "github.com/bxcodec/go-clean-arch/log/repository/mysql"
)

func TestStore(t *testing.T) {
	lg := &domain.Log{
		Id:         1,
		Pagination: 1,
		SearchWord: "badman",
		Response:   `{"Search":[{"imdbID":"tt0372784","Title":"Batman Begins","Year":"2005","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"},{"imdbID":"tt2975590","Title":"Batman v Superman: Dawn of Justice","Year":"2016","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"},{"imdbID":"tt0096895","Title":"Batman","Year":"1989","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg"},{"imdbID":"tt0103776","Title":"Batman Returns","Year":"1992","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOGZmYzVkMmItM2NiOS00MDI3LWI4ZWQtMTg0YWZkODRkMmViXkEyXkFqcGdeQXVyODY0NzcxNw@@._V1_SX300.jpg"},{"imdbID":"tt0118688","Title":"Batman \u0026 Robin","Year":"1997","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMGQ5YTM1NmMtYmIxYy00N2VmLWJhZTYtN2EwYTY3MWFhOTczXkEyXkFqcGdeQXVyNTA2NTI0MTY@._V1_SX300.jpg"},{"imdbID":"tt0112462","Title":"Batman Forever","Year":"1995","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNDdjYmFiYWEtYzBhZS00YTZkLWFlODgtY2I5MDE0NzZmMDljXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg"},{"imdbID":"tt4116284","Title":"The Lego Batman Movie","Year":"2017","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTcyNTEyOTY0M15BMl5BanBnXkFtZTgwOTAyNzU3MDI@._V1_SX300.jpg"},{"imdbID":"tt0103359","Title":"Batman: The Animated Series","Year":"1992–1995","Type":"series","Poster":"https://m.media-amazon.com/images/M/MV5BOTM3MTRkZjQtYjBkMy00YWE1LTkxOTQtNDQyNGY0YjYzNzAzXkEyXkFqcGdeQXVyOTgwMzk1MTA@._V1_SX300.jpg"},{"imdbID":"tt1569923","Title":"Batman: Under the Red Hood","Year":"2010","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"},{"imdbID":"tt2313197","Title":"Batman: The Dark Knight Returns, Part 1","Year":"2012","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"}],"totalResults":"445","Response":"True","Error":null}`,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := regexp.QuoteMeta(`INSERT log 
	SET 
		pagination=?,
		searchword=?,
		response=?, 
		created_at=?`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(lg.Pagination, lg.SearchWord, lg.Response, time.Now().Format("2006.01.02 15:04:05")).
		WillReturnResult(sqlmock.NewResult(1, 1))

	a := mysqlLogRepository.NewMysqlLogRepository(db)

	err = a.Store(context.TODO(), *lg)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), lg.Id)
}
