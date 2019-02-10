package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Endpoint struct {
	//  `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	gorm.Model
	Hits int64          `gorm:"default:0"`
	User string         `gorm:"-" json:"user" form:"user" query:"user"`
	Slug string         `gorm:"unique,not null" json:"slug" form:"slug" query:"slug"`
	Data postgres.Jsonb `gorm:"not null" json:"data" form:"data" query:"data"`
}
