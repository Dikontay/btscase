package repository

import (
	"context"

	"github.com/Dikontay/btscase/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) *Repository {
	return &Repository{conn}
}

func (r *Repository) GetOffersByMarket(ctx context.Context, market string) ([]models.Offer, error) {
	var offers []models.Offer
	result := r.conn.WithContext(ctx).Order("precent").Where("market = ?", market).Find(&offers)
	if result.Error != nil {
		return nil, result.Error
	}
	return offers, nil
}
