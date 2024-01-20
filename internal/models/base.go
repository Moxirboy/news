package models

import "time"

type By struct{
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
type At struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}