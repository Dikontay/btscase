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

func (s *Service) GetOffersByMarket(ctx context.Context, market string, userid int) ([]models.Offer, error) {
	return s.repo.GetOffersByMarket(ctx, market, userid)
}

func (s *Service) AddOffer(ctx context.Context, offer *models.Offer) error {
	return s.repo.AddOffer(ctx, offer)
}

func (s *Service) AddCard(ctx context.Context, card *models.Card) error {
	return s.repo.AddCard(ctx, card)
}

func (s *Service) AddOffers(ctx context.Context, offers []*models.Offer) error {

	return s.repo.AddOffers(ctx, offers)
}

func (s *Service) GetUserInfo(ctx context.Context, id int) (*models.User, error) {

	user, err := s.repo.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	user.CleanPassword()
	return user, err
}

func (s *Service) GetCardByUserID(ctx context.Context, userid int) (*models.Card, error) {
	return s.repo.GetCardByUserID(ctx, userid)
}
