package service

import (
	"myproject/forum/models"
	"myproject/forum/repository"
)

type IUserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uint64) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint64) error
	ListUsers() ([]models.User, error)
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.UserRepository.Create(user)
}

func (s *UserService) GetUserByID(id uint64) (*models.User, error) {
	return s.UserRepository.FindByID(id)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.UserRepository.FindByUsername(username)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.UserRepository.FindByEmail(email)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.UserRepository.Update(user)
}

func (s *UserService) DeleteUser(id uint64) error {
	return s.UserRepository.Delete(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.UserRepository.ListAll()
}
