package auth

import (
	"errors"
	"lgc/src/infraestructure/di"
	"lgc/src/infraestructure/middleware"
	"lgc/src/view/dto"
)

type LoginUseCase struct{}

func (uc *LoginUseCase) Execute(email, password string) (*dto.UserDTO, error) {

	userRepo := di.GetContainer().GetUserRepository()

	user, err := userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if !user.Exists() {
		return nil, errors.New("usuario no encontrado")
	}

	if !middleware.VerifyPassword(user.GetPassword(), password) {
		return nil, errors.New("contraseña incorrecta")
	}

	// Generar el token JWT con el secreto del usuario
	token, err := middleware.GenerateToken(user.GetID(), user.GetEmail())
	if err != nil {
		return nil, errors.New("error al generar el token")
	}

	userDTO := user.ToDTO()
	userDTO.SessionToken = token

	return userDTO, nil
}
