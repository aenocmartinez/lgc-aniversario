package usecase

import (
	"lgc/src/domain"
	"lgc/src/infraestructure/middleware"
	"lgc/src/view/dto"
)

type CrearUsuarioUseCase struct {
	usuaRepo domain.UserRepository
}

func NewCrearUsuarioUseCase(usuaRepo domain.UserRepository) *CrearUsuarioUseCase {
	return &CrearUsuarioUseCase{
		usuaRepo: usuaRepo,
	}
}

func (uc *CrearUsuarioUseCase) Execute(nombre, email, password string) dto.APIResponse {

	user, _ := uc.usuaRepo.FindByEmail(email)
	if user.Exists() {
		return dto.NewAPIResponse(209, "Existe un usuario con este correo electrónico", nil)
	}

	hashPassword, err := middleware.HashPassword(password)
	if err != nil {
		return dto.NewAPIResponse(500, err.Error(), nil)
	}

	user.SetName(nombre)
	user.SetEmail(email)
	user.SetPassword(hashPassword)

	err = uc.usuaRepo.Save(user)
	if err != nil {
		return dto.NewAPIResponse(500, err.Error(), nil)
	}

	return dto.NewAPIResponse(201, "El usuario se ha creado exitosamente", nil)
}
