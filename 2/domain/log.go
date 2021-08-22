package domain

import (
	"context"
)

// Log ...
type Log struct {
	Id         int64  `json:"id"`
	Pagination int64  `json:"pagination"`
	SearchWord string `json:"searchword"`
	Response   string `json:"response"`
}

// LogRepository ...
type LogRepository interface {
	Store(ctx context.Context, log Log) error
}
