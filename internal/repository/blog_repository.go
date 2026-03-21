package repository

import (
	"blog-api/internal/models"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) Create(blog *models.Blog) error {
	return r.db.Create(blog).Error
}

func (r *BlogRepository) GetByID(id uint) (*models.Blog, error) {
	var blog models.Blog

	if err := r.db.Preload("User").First(&blog, id).Error; err != nil {
		return nil, err
	}

	return &blog, nil

}

func (r *BlogRepository) GetAllPaginated(page, limit int, search string) ([]models.Blog, int64, error) {
	 
	var blogs []models.Blog
	var total int64

	query := r.db.Preload("User")

	// Apply search filter
	if search != "" {
		query = query.Where("title ILIKE ?", "%" + search+"%")
	}

	// Get total count 
	if err := query.Model(&models.Blog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page -1)*limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, 0, err
	}

	return blogs, total, nil

}

func (r *BlogRepository) GetByUserID(userID uint, page, limit int) ([]models.Blog, int64, error) {

	var blogs []models.Blog

	var total int64

	query := r.db.Where("user_id = ?", userID).Preload("User")

	// Get total count
	if err := query.Model(&models.Blog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page-1)*limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, 0, err
	}

	return blogs, total, nil

}

func (r *BlogRepository) Update(id uint, blog *models.Blog) error {
	return r.db.Model(&models.Blog{}, id).Updates(blog).Error
}

func (r *BlogRepository) Delete(id uint) error {
	return r.db.Delete(&models.Blog{}, id).Error
}

func (r *BlogRepository) Exists(id uint) bool {
	var count int64
	r.db.Model(&models.Blog{}).Where("id = ?", id).Count(&count)
	return count > 0
}
