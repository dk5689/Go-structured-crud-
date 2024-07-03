// service/user_service.go
package service

import (
	"crud_structured/model"
	"crud_structured/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) CreateUser(user *model.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.userRepository.Delete(id)
}
