package transport

import (
	"github.com/Dikontay/btscase/internal/service"
	"gorm.io/gorm"
)

type Transport struct {
	service *service.Service
}

func New(conn *gorm.DB) *Transport {
	return &Transport{service.New(conn)}
}
