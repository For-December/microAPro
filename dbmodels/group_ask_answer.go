package dbmodels

import "time"

type GroupAskAnswer struct {
	ID uint32 `gorm:"column:id;primaryKey;autoIncrement"`

	GroupID   string `gorm:"column:group_id;index;type:varchar(255)"`
	AskerID   string `gorm:"column:asker_id;type:varchar(255)"`
	ReplierID string `gorm:"column:replier_id;type:varchar(255)"`
	AskText   string `gorm:"column:ask_text;index;type:varchar(255)"`
	Answer    string `gorm:"column:answer;type:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
