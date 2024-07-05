// internal/models/recipe.go

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
	Reviews      []Review  `gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews"`
}

type RecipeRequest struct {
	Title        string                  `json:"title" validate:"required"`
	Description  string                  `json:"description"`
	Ingredients  string                  `json:"ingredients"`
	Instructions string                  `json:"instructions"`
	TagsIds      []uint                  `json:"tags_ids"`
	Images       []*multipart.FileHeader `json:"-"`
	ImageURLs    []string                `json:"-"`
}

type GetAllRecipeResponse struct {
	ID           uint    `gorm:"primaryKey"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Ingredients  string  `json:"ingredients"`
	Instructions string  `json:"instructions"`
	Images       []Image `gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"images"`
}

type DetailRecipeResponse struct {
	ID           uint             `gorm:"primaryKey"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Ingredients  string           `json:"ingredients"`
	Instructions string           `json:"instructions"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	Tags         []TagResponse    `gorm:"many2many:recipe_tags;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tags"`
	Images       []Image          `gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"images"`
	Reviews      []ReviewResponse `gorm:"foreignKey:RecipeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews"`
}
