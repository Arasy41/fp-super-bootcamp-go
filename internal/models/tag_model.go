package models

import "time"

type Tag struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique" json:"name"`
	Recipes   []Recipe  `gorm:"many2many:recipe_tags;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"recipes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TagRequest struct {
	Name string `json:"name" validate:"required"`
}
