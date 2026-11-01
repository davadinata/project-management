package services

import (
	"errors"

	"github.com/davadinata/project-management/models"
	"github.com/davadinata/project-management/repositories"
	"github.com/davadinata/project-management/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func newUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	existingUser, err := s.repo.FindByEmail(user.Email)

	if existingUser.InternalID != 0 {
		return errors.New("email already exists")
	}

	hased, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hased
	user.Role = "user"
	user.PublicID = uuid.New()
	return s.repo.Create(user)

}
