package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type ListarInscripcionesPendientesUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewListarInscripcionesPendientesUseCase(inscripcionRepo domain.InscripcionRepository) *ListarInscripcionesPendientesUseCase {
	return &ListarInscripcionesPendientesUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *ListarInscripcionesPendientesUseCase) Execute() dto.APIResponse {

	result := uc.inscripcionRepo.ListarInscripcionesPorEstado("PreAprobada")

	inscripciones := []dto.InscripcionDTO{}
	for _, inscripcion := range result {

		inscripciones = append(inscripciones, inscripcion.ToDTO())
	}

	return dto.NewAPIResponse(200, "Inscripciones pendientes", inscripciones)
}
