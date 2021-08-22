package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type mysqlLogRepository struct {
	Conn *sql.DB
}

// NewMysqlLogRepository ...
func NewMysqlLogRepository(Conn *sql.DB) domain.LogRepository {
	return &mysqlLogRepository{Conn}
}

func (m *mysqlLogRepository) Store(ctx context.Context, log domain.Log) (err error) {
	query := `INSERT log 
		SET 
			pagination=?,
			searchword=?,
			response=?, 
			created_at=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		log.Pagination,
		log.SearchWord,
		log.Response,
		time.Now().Format("2006.01.02 15:04:05"),
	)

	if err != nil {
		return err
	}

	return err
}
