package models

import (
	"github.com/davadinata/project-management/models/types"
	"github.com/google/uuid"
)

type ListPosition struct {
	InternaID int64           `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PubliicID uuid.UUID       `json:"public_id" db:"public_id" gorm:"public_id"`
	BoardID   int64           `json:"board_id" db:"board_id" gorm:"board_id"`
	ListOrder types.UUIDArray `json:"list_order"`
}
