package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type InscripcionesAprobadasUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewListarInscripcionesAprobadasUseCase(inscripcionRepo domain.InscripcionRepository) *InscripcionesAprobadasUseCase {
	return &InscripcionesAprobadasUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *InscripcionesAprobadasUseCase) Execute() dto.APIResponse {

	result := uc.inscripcionRepo.ListarInscripcionesPorEstado("Aprobada")

	inscripciones := []dto.InscripcionDTO{}
	for _, inscripcion := range result {

		inscripciones = append(inscripciones, inscripcion.ToDTO())
	}

	return dto.NewAPIResponse(200, "Inscripciones aprobadas", inscripciones)
}
