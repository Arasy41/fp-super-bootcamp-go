package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"size:255;not null" json:"username"`
	Email     string     `gorm:"size:255;unique;not null" json:"email" validate:"required,email"`
	Password  string     `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Profile   Profile    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"profile"`
	Reviews   []Review   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"reviews"`
	Favorites []Favorite `gorm:"many2many:user_favorites;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"favorites"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
