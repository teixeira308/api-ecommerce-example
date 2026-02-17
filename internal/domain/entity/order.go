package entity

import "time"

type Order struct {
	ID        string
	ItemID    string
	Quantity  int
	Total     float64
	CreatedAt time.Time
}
