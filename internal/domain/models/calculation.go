package domain

import "time"

type Calculation struct {
	ID        string
	Operation string
	A         int
	B         int
	Result    int
	CreatedAt time.Time
}
