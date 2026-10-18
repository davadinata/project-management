package models

type CardAssignee struct {
	InternalID int64 `json:"internal_id" db:"internal_id" gorm:"column:user_internal_id"`
	UserID     int64 `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id"`
}
