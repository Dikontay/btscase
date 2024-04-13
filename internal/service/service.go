package service

import (
	"context"

	"github.com/Dikontay/btscase/internal/models"
	"github.com/Dikontay/btscase/internal/repository"
	"gorm.io/gorm"
)

type Service struct {
	repo *repository.Repository
}

func New(conn *gorm.DB) *Service {
	return &Service{repository.New(conn)}
}

func (s *Service) GetOffersByMarket(ctx context.Context, market string) ([]models.Offer, error) {
	return s.repo.GetOffersByMarket(ctx, market)
}
