package models

import "time"

type Card struct {
	ID     int
	Bank   string
	Type   string
	Nomer  string
	Due    time.Time
	UserID int
}
