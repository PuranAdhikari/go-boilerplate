package service

import (
	"errors"
	"github.com/puranadhikari/go-boilerplate/internal/dto"
	"github.com/puranadhikari/go-boilerplate/internal/model"
	"github.com/puranadhikari/go-boilerplate/internal/repository"
)

type UserService struct {
	Ur repository.IUserRepository
}

func (s *UserService) Register(email, password string) (*dto.RegisterResponse, error) {
	exists, err := s.Ur.EmailExists(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already in use")
	}
	// Hash the password (using bcrypt or similar)
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = s.Ur.CreateUser(user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	return response, err
}

func hashPassword(password string) (string, error) {
	// Use bcrypt or another library to hash the password
	// This is just a placeholder function
	return password, nil
}
