package models

import "github.com/google/uuid"

type City struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
}
