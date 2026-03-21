package repository

import (
	"blog-api/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {

	var user models.User

	if err := r.db.Preload("Blogs").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
 
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User

	if err := r.db.Find(&users).Error; err !=nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(id uint, user *models.User) error {
	return r.db.Model(&models.User{}, id).Updates(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}