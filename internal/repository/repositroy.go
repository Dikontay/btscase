package repository

import "database/sql"

type Repository struct {
	conn *sql.Conn
}

func New(conn *sql.Conn) *Repository {
	return &Repository{conn}
}
