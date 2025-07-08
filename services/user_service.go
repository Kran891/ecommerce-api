package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repository"
	"ecommerce-api/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *models.User) error
	LoginUser(email, password string) (string, error)
	Find(id uuid.UUID) (*models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
	CartItems(id uuid.UUID) (user *models.User, err error)
}
type userService struct {
	repo repository.UserRepository
}

// CartItems implements UserService.
func (s *userService) CartItems(id uuid.UUID) (user *models.User, err error) {
	return s.repo.CartItems(id)
}

// Delete implements UserService.
func (s *userService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

// Update implements UserService.
func (s *userService) Update(user *models.User) error {
	dbUser, err := s.repo.FindByID(user.ID)
	if err != nil {
		return err
	}
	dbUser.Username = user.Username
	dbUser.Email = user.Email
	dbUser.UpdatedAt = time.Now()
	err = s.repo.Update(dbUser)
	if err != nil {
		return err
	}
	return nil
}

// Find implements UserService.
func (s *userService) Find(id uuid.UUID) (*models.User, error) {
	return s.repo.FindByID(id)
}
func (s *userService) RegisterUser(user *models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.repo.Create(user)
}

func (s *userService) LoginUser(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	return utils.MakeJWT(user)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
