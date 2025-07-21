package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type AnularInscripcionUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewAnularInscripcionUseCase(inscripcionRepo domain.InscripcionRepository) *AnularInscripcionUseCase {
	return &AnularInscripcionUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *AnularInscripcionUseCase) Execute(inscripcionID int64) dto.APIResponse {

	inscripcion := uc.inscripcionRepo.BuscarPorID(inscripcionID)
	if !inscripcion.Existe() {
		return dto.NewAPIResponse(404, "Inscripción no encontrada", nil)
	}

	if inscripcion.EstaRechazada() {
		return dto.NewAPIResponse(200, "La inscripción ya ha sido rechazada.", nil)
	}

	exito := inscripcion.Anular()

	if !exito {
		return dto.NewAPIResponse(500, "Ha ocurrido un error interno en el sistema. Por favor, intenta nuevamente más tarde.", nil)
	}

	return dto.NewAPIResponse(200, "Inscripción ha sido anulada exitosamente", nil)
}
