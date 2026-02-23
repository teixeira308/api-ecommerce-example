package entity

import "time"

type Order struct {
	ID        string
	Items     []*OrderItem
	Total     float64
	Status    string
	CreatedAt time.Time
}
