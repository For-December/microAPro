package dbmodels

import "time"

type GroupLog struct {
	ID uint32 `gorm:"column:id;primaryKey;autoIncrement"`

	GroupID string `gorm:"column:group_id;index;type:varchar(255)"`
	UserID  string `gorm:"column:user_id;type:varchar(255)"`

	Message string `gorm:"column:message;type:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
