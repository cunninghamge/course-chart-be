package models

import "time"

type ModuleActivity struct {
	ID                   int       `json:"id"`
	Input                int       `json:"input"`
	Notes                string    `json:"notes"`
	ModuleId             int       `json:"-"`
	CreatedAt, UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	ActivityId           int       `json:"activityId"`
	Activity             Activity  `json:"activity"`
}
