package testemodel

import "time"

type TesteModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
	