package models

import "github.com/google/uuid"

type Vendor struct {
	ID uuid.UUID `json:"id" gorm:"typw:uuid;default:uuid_generate_v4();prinary_key"`
}
