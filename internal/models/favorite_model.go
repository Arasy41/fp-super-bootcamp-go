package models

import "time"

type Favorite struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	RecipeID  uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"recipe_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FavoriteRequest struct {
	UserID   uint `json:"user_id" validate:"required"`
	RecipeID uint `json:"recipe_id" validate:"required"`
}
