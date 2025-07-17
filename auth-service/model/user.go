package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	MSISDN   string    `gorm:"uniqueIndex;not null" json:"msisdn"`
	Username string    `gorm:"uniqueIndex;not null" json:"username"`
	Name     string    `gorm:"not null" json:"name"`
	Password string    `gorm:"not null" json:"-"`
	Role     string    `gorm:"default:user" json:"role"`
}
