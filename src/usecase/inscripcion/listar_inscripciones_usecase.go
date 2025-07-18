package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type ListarInscripcionesUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewListarInscripcionesUseCase(inscripcionRepo domain.InscripcionRepository) *ListarInscripcionesUseCase {
	return &ListarInscripcionesUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *ListarInscripcionesUseCase) Execute() dto.APIResponse {

	result := uc.inscripcionRepo.Listar()

	inscripciones := []dto.InscripcionDTO{}
	for _, inscripcion := range result {

		inscripciones = append(inscripciones, inscripcion.ToDTO())
	}

	return dto.NewAPIResponse(200, "Listado de Inscripciones", inscripciones)
}
