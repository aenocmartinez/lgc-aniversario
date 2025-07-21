package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type ListarInscripcionesUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewListarInscripcionesUseCase(repo domain.InscripcionRepository) *ListarInscripcionesUseCase {
	return &ListarInscripcionesUseCase{inscripcionRepo: repo}
}

func (uc *ListarInscripcionesUseCase) Execute() dto.APIResponse {

	resultado := uc.inscripcionRepo.ListarConParticipantes()

	return dto.NewAPIResponse(200, "Listado de inscripciones con participantes", resultado)
}
