package usecase

import (
	"lgc/src/domain"
	"lgc/src/infraestructure/middleware"
	"lgc/src/view/dto"
	formrequest "lgc/src/view/form-request"
)

type ActualizarUsuarioUseCase struct {
	userRepository domain.UserRepository
}

func NewActualizarUsuarioUseCase(userRepository domain.UserRepository) *ActualizarUsuarioUseCase {
	return &ActualizarUsuarioUseCase{
		userRepository: userRepository,
	}
}

func (uc *ActualizarUsuarioUseCase) Executar(req formrequest.ActualizarUsuarioFormRequest) dto.APIResponse {

	user, _ := uc.userRepository.FindByID(req.ID)
	if !user.Exists() {
		return dto.NewAPIResponse(404, "Usuario no encontrado", nil)
	}

	hashPassword, err := middleware.HashPassword(req.Password)
	if err != nil {
		return dto.NewAPIResponse(500, err.Error(), nil)
	}

	user.SetName(req.Nombre)
	user.SetEmail(req.Email)
	user.SetPassword(hashPassword)

	if err := user.Update(); err != nil {
		return dto.NewAPIResponse(500, "Ha ocurrido un error en el sistema", nil)
	}

	return dto.NewAPIResponse(200, "Usuario actualizado con éxito", user.ToDTO())
}
