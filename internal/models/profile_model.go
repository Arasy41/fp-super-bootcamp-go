package models

import "time"

type Profile struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"unique;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	FullName  string    `json:"full_name"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProfileRequest struct {
	FullName  string `json:"full_name"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}
