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
	result := r.conn.WithContext(ctx).
		Joins("JOIN cards ON cards.bank = offers.bank").
		Joins("JOIN users ON users.id = cards.userid").
		Where("offers.market = ?", market).
		Order("offers.precent").
		Find(&offers)
	if result.Error != nil {
		return nil, result.Error
	}
	return offers, nil
}

func (r *Repository) AddOffer(ctx context.Context, offer *models.Offer) error {
	result := r.conn.WithContext(ctx).Create(offer)
	return result.Error
}

func (r *Repository) AddCard(ctx context.Context, card *models.Card) error {
	result := r.conn.WithContext(ctx).Create(card)
	return result.Error
}

func (r *Repository) AddOffers(ctx context.Context, offers []*models.Offer) error {
	result := r.conn.WithContext(ctx).Create(offers)
	return result.Error
}

func (r *Repository) GetUserInfo()
