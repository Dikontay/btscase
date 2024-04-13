package service

import (
	"github.com/Dikontay/btscase/internal/repository"
	"gorm.io/gorm"
)

type Service struct {
	repo *repository.Repository
}

func New(conn *gorm.DB) *Service {
	return &Service{repository.New(conn)}
}
