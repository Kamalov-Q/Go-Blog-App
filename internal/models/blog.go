package models

import "time"

type Blog struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"type:text;not null"`
	UserID uint `json:"user_id" gorm:"not null"`
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;constaint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}

type CreateBlogRequest struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
}

type UpdateBlogRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type PaginationParams struct {
	Page int `form:"page, default=1"`
	Limit int `form:"limit,default=10"`
	Search string `form:"search"`

}