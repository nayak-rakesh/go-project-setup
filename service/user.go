package service

import "github.com/nayak-rakesh/go-project-setup/domain"

// UserService ...
type UserService struct {
	userRepo domain.UserRepository
}

// NewUserService ...
func NewUserService(u domain.UserRepository) domain.UserSerice {
	return &UserService{
		userRepo: u,
	}
}

// Create ...
func (service *UserService) Create(u *domain.User) error {
	err := service.userRepo.Create(u)
	if err != nil {
		return err
	}
	return nil
}

// GetByEmail ...
func (service *UserService) GetByEmail(email string) (domain.User, error) {
	user, err := service.userRepo.GetByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// GetAll ...
func (service *UserService) GetAll() ([]domain.User, error) {
	users, err := service.userRepo.GetAll()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}
