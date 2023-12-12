package user

import (
	"errors"
	"layanan-kependudukan-api/helper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetUsers(pagination helper.Pagination) (helper.Pagination, error)
	RegiserUser(input CreateUserInput) (User, error)
	UpdateUser(ID GetUserDetailInput, input CreateUserInput) (User, error)
	Login(input LoginInput) (User, error)
	Logout(user User) (User, error)
	IsEmailAvailable(input EmailInput) (bool, error)
	UpdateToken(user User, token string) (User, error)
	GetUserById(ID int) (User, error)
	GetUserByNIK(NIK string) (User, error)
	GetUserByAdmin() ([]User, error)
	DeleteUser(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) RegiserUser(input CreateUserInput) (User, error) {

	user := User{}
	user.Nik = input.Nik
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.AvatarPath = input.AvatarPath
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.Role = input.Role
	user.Token = input.Token

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Email or password wrong!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, errors.New("Email or password wrong!")
	}

	if input.Type == "WEB" {
		if user.Role == "PENDUDUK" || user.Role == "RT/RW" {
			return user, errors.New("User yang dipakai tidak dapat login pada website")
		}
	}

	if input.Token != "" {
		user, err = s.UpdateToken(user, input.Token)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (s *service) Logout(user User) (User, error) {
	user.Token = ""
	user, err := s.repository.Update(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input EmailInput) (bool, error) {

	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID != 0 {
		return false, errors.New("Email already taken!")
	}

	return true, nil
}

func (s *service) GetUserById(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found!")
	}

	return user, nil
}

func (s *service) GetUserByNIK(ID string) (User, error) {
	user, err := s.repository.FindByNIK(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found!")
	}

	return user, nil
}

func (s *service) GetUserByAdmin() ([]User, error) {
	user, err := s.repository.FindByAdmin()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUsers(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}

func (s *service) UpdateToken(currentUser User, token string) (User, error) {
	user := currentUser

	user.Token = token

	user.UpdatedAt = time.Now()

	newUser, err := s.repository.Update(user)
	return newUser, err
}

func (s *service) UpdateUser(inputDetail GetUserDetailInput, input CreateUserInput) (User, error) {
	user, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return user, err
	}

	user.Nik = input.Nik
	user.Name = input.Name
	user.Email = input.Email
	if input.Password != "" {
		user.Password = input.Password
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

		if err != nil {
			return user, err
		}

		user.Password = string(password)
	}

	user.Role = input.Role

	user.AvatarPath = input.AvatarPath

	user.UpdatedAt = time.Now()

	newUser, err := s.repository.Update(user)
	return newUser, err
}

func (s *service) DeleteUser(ID int) error {
	user, errId := s.repository.FindByID(ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(user)
	return err
}
