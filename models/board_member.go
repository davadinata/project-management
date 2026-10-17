package models

type BoardMember struct {
	BoardID  int64 `json:"board_internal_id" db:"board_internal_id" gorm:"column:board_internal_id;primaryKey"`
	UserID   int64 `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id;primaryKey"`
	JoinedAt int64 `json:"joined_at" db:"joined_at"`
}
