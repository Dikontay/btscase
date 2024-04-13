package repository

import "gorm.io/gorm"

type Repository struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) *Repository {
	return &Repository{conn}
}
