package transport

import "github.com/RickDred/ozinse/internal/auth"

type transport struct {
	service auth.ServiceInterface
}

func New(s auth.ServiceInterface) auth.HandlersInterface {
	return &transport{service: s}
}
