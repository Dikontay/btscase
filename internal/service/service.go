package service

import (
	"database/sql"

	"github.com/Dikontay/btscase/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(conn *sql.Conn) *Service {
	return &Service{repository.New(conn)}
}
