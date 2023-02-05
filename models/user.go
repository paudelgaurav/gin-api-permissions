package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string         `gorm:"size:255;not null;unique" json:"username" form:"username" binding:"required"`
	Permissions pq.StringArray `gorm:"type:varchar(200)[]" json:"permissions" form:"permissions"`
}
