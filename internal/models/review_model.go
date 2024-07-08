package models

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	RecipeID  uint      `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"recipe_id"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID" json:"user" swaggerignore:"true"`
	Recipe    Recipe    `gorm:"foreignKey:RecipeID" json:"review" swaggerignore:"true"`
}

type ReviewRequest struct {
	UserID   uint   `json:"user_id" validate:"required"`
	RecipeID uint   `json:"recipe_id" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

type ReviewResponse struct {
	ID        uint         `gorm:"primaryKey"`
	UserID    uint         `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	RecipeID  uint         `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"recipe_id"`
	Content   string       `gorm:"type:text" json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserResponse `gorm:"foreignKey:UserID" json:"user"`
	Recipe    Recipe       `gorm:"foreignKey:RecipeID" json:"recipe"`
}
