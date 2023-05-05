package service

import (
	cliente "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/client"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/dto"
	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/errors"
	funciones "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/funciones"
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/model"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.User, e.ApiError)
	GetUsers() (dto.Users, e.ApiError)
	NewUser(userDto dto.User) (dto.User, e.ApiError)
	DeleteUser(id int) e.ApiError
}

var jwtKey = []byte("secret_key")

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.User, e.ApiError) {

	var user model.User = cliente.GetUserById(id)
	var userDto dto.User

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.Id = user.Id
	userDto.Username = user.Username
	userDto.Password = user.Password
	userDto.Email = user.Email
	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName

	return userDto, nil
}

func (s *userService) GetUsers() (dto.Users, e.ApiError) {

	var users model.Users = cliente.GetUsers()
	var usersDto dto.Users

	for _, user := range users {
		var userDto dto.User
		userDto.Id = user.Id
		userDto.Username = user.Username
		userDto.Password = user.Password
		userDto.Email = user.Email
		userDto.FirstName = user.FirstName
		userDto.LastName = user.LastName
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) NewUser(userDto dto.User) (dto.User, e.ApiError) {
	var user model.User

	user.Username = userDto.Username
	user.Email = userDto.Email
	user.Password = funciones.SSHA256(userDto.Password)
	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName

	var err e.ApiError
	user, err = cliente.NewUser(user)

	userDto.Password = user.Password
	userDto.Id = user.Id

	return userDto, err

}

func (s *userService) DeleteUser(id int) e.ApiError {
	err := cliente.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}
