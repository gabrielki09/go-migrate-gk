package estemodel

import "time"

type EsteModel struct {
	UUID      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
	