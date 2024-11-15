package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type User struct {
	Id        ulid.ULID `json:"id" gorm:"primaryKey;not null"`
	Name      string    `json:"name" form:"name" validate:"required" gorm:"not null"`
	Username  string    `json:"username" form:"username" validate:"required,gte=6,lte=32" gorm:"unique;not null"`
	Password  string    `json:"password" form:"password" validate:"required,gte=8"`
	Email     string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Picture   string    `json:"picture" form:"picture"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoUpdateTime;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoCreateTime;not null"`
}
