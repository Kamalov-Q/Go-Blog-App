package models

import "time"

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Blogs []Blog `json:"blogs,omitempty" gorm:"foreignKey:UserID; containt:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required, email"`
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required, email"`
}

