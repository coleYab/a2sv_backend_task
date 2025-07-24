package usecases

import (
	"task_manager/delivery/controllers/dto"
	"task_manager/domain"
	"task_manager/repositories"
)

type IUserUseCase interface {
	CreateUser(userPayload dto.RegisterUserDto, userRole string) (domain.User, error)
	FindUserByUserName(userName string) (domain.User, error)
	PromoteUser(userID string) (domain.User, error)
	FindUserByID(userID string) (domain.User, error)
	GetUserCount() int64
}

type UserUseCase struct {
	userRepository repositories.IUserRepository
}

func NewUserUseCase(repo repositories.IUserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: repo,
	}
}

func (us *UserUseCase) PromoteUser(userID string) (domain.User, error) {
	return us.userRepository.PromoteUser(userID)
}

func (us *UserUseCase) CreateUser(userPayload dto.RegisterUserDto, userRole string) (domain.User, error) {
	return us.userRepository.CreateUser(userPayload, userRole)
}

func (us *UserUseCase) GetUserCount() int64 {
	return us.userRepository.GetUserCount()
}

func (us *UserUseCase) FindUserByID(userID string) (domain.User, error) {
	return us.userRepository.FindUserByID(userID)
}

func (us *UserUseCase) FindUserByUserName(UserName string) (domain.User, error) {
	return us.userRepository.FindUserByUserName(UserName)
}
