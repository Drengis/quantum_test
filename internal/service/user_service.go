package service

import (
	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/models"
	"github.com/user/quantum-server/internal/repository"
)

type UserService interface {
	FindOrCreate(dto dto.CreateUserDto) (*models.User, error)
	FindByTgID(tgID string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) FindByTgID(tgID string) (*models.User, error) {
	return s.repo.FindByTgID(tgID)
}

func (s *userService) FindOrCreate(userDto dto.CreateUserDto) (*models.User, error) {
	user, err := s.repo.FindByTgID(userDto.TgID)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	newUser := &models.User{
		TgID:      userDto.TgID,
		Username:  getStringValue(userDto.Username),
		FirstName: getStringValue(userDto.FirstName),
		LastName:  getStringValue(userDto.LastName),
		LangCode:  getStringValue(userDto.LangCode),
		InvitedBy: getStringValue(userDto.InvitedBy),
		IsActive:  true,
	}

	err = s.repo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
