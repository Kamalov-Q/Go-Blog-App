package service

import (
	"blog-api/internal/models"
	"blog-api/internal/repository"
	"fmt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {

	// Validate email uniqueness
	existingUser, _ := s.repo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("Email already exists!")
	}

	user := &models.User{
		Name: req.Name,
		Email: req.Email,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("User not found!")
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) UpdateUser(id uint, req *models.UpdateUserRequest) (*models.User, error) {

	// Checking if the user exists
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("User not found!")
	}

	// Checking if the emai is unique if email is being updated
	if req.Email != "" && req.Email != user.Email {
		existingUser, _ := s.repo.GetByEmail(req.Email)

		if existingUser != nil {
			return nil, fmt.Errorf("Email already exists!")
		}
		user.Email = req.Email

	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if err := s.repo.Update(id, user); err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) DeleteUser(id uint) error {
	// Checking if the user exists

	if _, err := s.repo.GetByID(id); err != nil {
		return fmt.Errorf("User not found!")
	}

	return s.repo.Delete(id)

}
