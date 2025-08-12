// Package usecases
package usecases

import (
	"blogger/modules/user/application/dto"
	"blogger/modules/user/domain/entity"
	"blogger/modules/user/domain/repository"
	"blogger/modules/utils"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserUseCase struct {
	passwordUtil utils.PasswordUtils 
	tokenUtil utils.AuthToken
	tokenRepostiory repository.TokenRepository
	userRepository repository.UserRepository
}

func NewUserUseCase(
	userRepository repository.UserRepository,
	tokenRepository repository.TokenRepository,
	passwordUtil utils.PasswordUtils,
	tokenUtil utils.AuthToken,
	) *UserUseCase {
	uc := &UserUseCase{
		userRepository: userRepository,
		tokenRepostiory: tokenRepository,
		tokenUtil: tokenUtil,
		passwordUtil: passwordUtil,
	}

	return uc
}

func (uc *UserUseCase) UpdateProfile(userId string, updateProfileDto dto.UpdateProfileDTO) (entity.User, error) {
	user, err := uc.userRepository.GetUserByID(userId)
	if err != nil {
		return entity.User{}, err 
	}

	user.UpdateProfile(updateProfileDto.Bio, updateProfileDto.ProfilePicture)
	err = uc.userRepository.UpdateUser(user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (uc *UserUseCase) CreateUser(registerUserDto dto.RegisterUserDTO) (entity.User, error) {
	hashedPassword := uc.passwordUtil.HashPassword(registerUserDto.Password)
	if err := uc.checkIfUnique(registerUserDto); err != nil {
		return entity.User{}, err
	}

	user := entity.NewUser(
		uuid.NewString(), registerUserDto.Username, 
		registerUserDto.Email, "user", hashedPassword, time.Now(), 
		time.Now(),
	)
	err := uc.userRepository.CreateUser(user)
	return user, err
}	

func (uc *UserUseCase) Login(loginDto dto.LoginUserDTO) (utils.Token, error) {
	user, err := uc.userRepository.GetUserByUsername(loginDto.Username)
	log.Println(user, err)
	if err != nil {
		return utils.Token{}, fmt.Errorf("invalid credentials")
	}

	log.Println(user.Password)

	if err := uc.passwordUtil.ComparePassword(user.Password, loginDto.Password); err != nil {
		return utils.Token{}, fmt.Errorf("invalid credentials")
	}

	return uc.generateToken(user)
}

func (uc *UserUseCase) generateToken(user entity.User) (utils.Token, error) {
	accessToken, err := uc.tokenUtil.Generate(user.ID, user.Username)
	if err != nil {
		return utils.Token{}, err 
	}

	tokenString := uuid.NewString()
	expiresAfter := time.Now().Add(time.Hour * 24)
	refreshToken := entity.NewToken(uuid.NewString(), user.ID, uc.passwordUtil.HashPassword(tokenString), expiresAfter, time.Now(), false)
	if err := uc.tokenRepostiory.CreateToken(refreshToken); err != nil {
		return utils.Token{}, err 
	}

	return utils.Token{AccessToken: accessToken, RefreshToken: tokenString}, nil	
}

func (uc *UserUseCase) checkIfUnique(registerUserDto dto.RegisterUserDTO) error {
	userExsists := fmt.Errorf("the user is not unique")
	fmt.Println(uc.userRepository == nil)
	if _, err := uc.userRepository.GetUserByUsername(registerUserDto.Username); err == nil {
		return userExsists
	}

	if _, err := uc.userRepository.GetUserByEmail(registerUserDto.Email); err == nil {
		return userExsists
	}

	return nil
}

func (uc *UserUseCase) PromoteUser(promoterID string, promoteUserDto dto.PromoteUserDTO) error {
	promoter, err := uc.userRepository.GetUserByID(promoterID)
	if err != nil || promoter.Role != "admin" {
		return fmt.Errorf("invalid resource")
	}

	promotee, err := uc.userRepository.GetUserByID(promoteUserDto.ID)
	if err != nil {
		return fmt.Errorf("promotee not found")
	}

	promotee.Promote("admin")
	return  uc.userRepository.UpdateUser(promotee)
}