package transport

import (
	"database/sql"

	"github.com/Dikontay/btscase/internal/service"
)

type Transport struct {
	service *service.Service
}

func New(conn *sql.Conn) *Transport {
	return &Transport{service.New(conn)}
}
