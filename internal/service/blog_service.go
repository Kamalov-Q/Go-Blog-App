package service

import (
	"blog-api/internal/models"
	"blog-api/internal/repository"
	"fmt"
)

type BlogService struct {
	blogRepo *repository.BlogRepository
	userRepo *repository.UserRepository
}

func NewBlogService(blogRepo *repository.BlogRepository, userRepo *repository.UserRepository) *BlogService {
	return &BlogService{
		blogRepo: blogRepo,
		userRepo: userRepo,
	}
}

func (s *BlogService) CreateBlog(req *models.CreateBlogRequest) (*models.Blog, error) {
	// Verify if the user exists
	_, err := s.blogRepo.GetByID(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("User not found!")
	}

	blog := &models.Blog{
		Title: req.Title,
		Content: req.Content,
		UserID: req.UserID,
	}

	if err := s.blogRepo.Create(blog); err != nil {
		return nil, err
	}

	// Reloading with user data
	return s.blogRepo.GetByID(blog.ID)
}

func (s *BlogService) GetBlogByID(id uint) (*models.Blog, error) {
	blog, err := s.blogRepo.GetByID(id)

	if err != nil {
		return nil, fmt.Errorf("Blog not found!")
	}

	return blog, nil

}

func (s *BlogService) GetAllBlogsPaginated(page, limit int, search string) ([]models.Blog, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 || limit > 100 {
		limit = 10
	}

	return s.blogRepo.GetAllPaginated(page, limit, search)
}

func (s *BlogService) UpdateBlog(id uint, req *models.UpdateBlogRequest) (*models.Blog, error) {

	// Check if the blog exists
	blog, err := s.blogRepo.GetByID(id)

	if err != nil {
		return nil, fmt.Errorf("Blog not found!")
	}

	if req.Title != "" {
		blog.Title = req.Title
	}

	if req.Content != "" {
		blog.Content = req.Content
	}

	if err := s.blogRepo.Update(id, blog); err != nil {
		return nil, err
	}

	return s.blogRepo.GetByID(id)
}

func (s *BlogService) DeleteBlog(id uint) error {
	if !s.blogRepo.Exists(id) {
		return fmt.Errorf("Blog not found!")
	}

	return s.blogRepo.Delete(id)
}

