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

func (r *Repository) GetOffersByMarket(ctx context.Context, market string, userid int) ([]models.Offer, error) {
	var offers []models.Offer
	result := r.conn.WithContext(ctx).
		Joins("JOIN cards ON cards.bank = offers.bank").
		Joins("JOIN users ON users.id = cards.userid").
		Where("offers.market = ?", market).
		Where("users.id = ?", userid).
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

func (r *Repository) GetUserInfo(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	result := r.conn.WithContext(ctx).Where(`id=?`, id).Find(&user)
	return &user, result.Error
}

func (r *Repository) GetCardByUserID(ctx context.Context, id int) (*models.Card, error) {
	var card models.Card
	result := r.conn.WithContext(ctx).Where(`userid=?`, id).Find(&card)
	return &card, result.Error
}
