package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type BuscarInscripcionUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewBuscarInscripcionUseCase(inscripcionRepo domain.InscripcionRepository) *BuscarInscripcionUseCase {
	return &BuscarInscripcionUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *BuscarInscripcionUseCase) Execute(inscripcionID int64) dto.APIResponse {

	inscripcion := uc.inscripcionRepo.BuscarPorID(inscripcionID)
	if !inscripcion.Existe() {
		return dto.NewAPIResponse(404, "Inscripción no encontrada", nil)
	}

	return dto.NewAPIResponse(200, "Inscripción encontrada", inscripcion.ToDTO())
}
