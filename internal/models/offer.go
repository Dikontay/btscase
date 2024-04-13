package models

import "time"

type Offer struct {
	ID         int
	Bank       string
	Category   string
	Precent    float64
	Condition  string
	Due        time.Time
	Limitation string
}
