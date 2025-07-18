package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type ListarInscripcionesValidadasUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewListarInscripcionesValidadasUseCase(inscripcionRepo domain.InscripcionRepository) *ListarInscripcionesValidadasUseCase {
	return &ListarInscripcionesValidadasUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *ListarInscripcionesValidadasUseCase) Execute() dto.APIResponse {

	result := uc.inscripcionRepo.ListarInscripcionesPorEstado("Validado")

	inscripciones := []dto.InscripcionDTO{}
	for _, inscripcion := range result {

		inscripciones = append(inscripciones, inscripcion.ToDTO())
	}

	return dto.NewAPIResponse(200, "Inscripciones validadas", inscripciones)
}
