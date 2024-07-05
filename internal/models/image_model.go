package models

import "time"

type Image struct {
	ID        uint      `gorm:"primaryKey"`
	RecipeID  uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"recipe_id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ImageRequest struct {
	RecipeID uint   `json:"recipe_id" validate:"required"`
	URL      string `json:"url" validate:"required"`
}
