package models

import (
	"mime/multipart"
	"time"
)

type Recipe struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Ingredients  string    `json:"ingredients"`
	Instructions string    `json:"instructions"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Tags         []Tag     `gorm:"many2many:recipe_tags;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tags"`
	Images       []Image   `gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"images"`
	Reviews      []Review  `gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews" swaggerignore:"true"`
}

type RecipeRequest struct {
	Title        string                  `json:"title"`
	Description  string                  `json:"description"`
	Ingredients  string                  `json:"ingredients"`
	Instructions string                  `json:"instructions"`
	Images       []*multipart.FileHeader `json:"images"`
	ImageURLs    []string                `json:"image_urls"`
	TagIDs       []uint                  `json:"tag_ids"`
}

type RecipeTag struct {
	RecipeID uint `gorm:"index"`
	TagID    uint `gorm:"index"`
}
