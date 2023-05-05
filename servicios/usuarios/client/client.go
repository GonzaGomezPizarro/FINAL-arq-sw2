package client

import (
	"errors"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/model"

	e "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/usuarios/errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User
	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)
	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)
	log.Debug("Users: ", users)
	return users
}

func NewUser(user model.User) (model.User, e.ApiError) {

	var existingUser model.User
	err := Db.Where("username = ?", user.Username).First(&existingUser).Error
	if err == nil {
		// El nombre de usuario ya está en uso, devolver un error al cliente
		return existingUser, e.NewBadRequestApiError("username in use")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Se produjo un error al buscar el usuario en la base de datos
		return model.User{}, e.NewInternalServerApiError("username cannot be checked", err)
	}
	// El nombre de usuario no está en uso, revisemos el mail
	errr := Db.Where("email = ?", user.Email).First(&existingUser).Error
	if errr == nil {
		// El nombre de usuario ya está en uso, devolver un error al cliente
		return existingUser, e.NewBadRequestApiError("email in use")
	} else if !errors.Is(errr, gorm.ErrRecordNotFound) {
		// Se produjo un error al buscar el email en la base de datos
		return model.User{}, e.NewInternalServerApiError("email cannot be checked", err)
	}
	//crear un nuevo registro en la base de datos
	result := Db.Create(&user)
	if result.Error != nil {
		log.Error("Error al crear el usuario:", result.Error)
		return model.User{}, e.NewInternalServerApiError("user cannot be created", err)
	}
	log.Debug("User Created: ", user)
	return user, nil

}

func DeleteUser(userID int) e.ApiError {
	var user model.User
	result := Db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return e.NewNotFoundApiError("user not found")
		} else {
			return e.NewInternalServerApiError("error while searching for user", result.Error)
		}
	}
	result = Db.Delete(&user)
	if result.Error != nil {
		return e.NewInternalServerApiError("error while deleting user", result.Error)
	}
	return nil
}
