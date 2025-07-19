package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type AprobarInscripcionUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewAprobarInscripcionUseCase(inscripcionRepo domain.InscripcionRepository) *AprobarInscripcionUseCase {
	return &AprobarInscripcionUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *AprobarInscripcionUseCase) Execute(inscripcionID int64) dto.APIResponse {

	inscripcion := uc.inscripcionRepo.BuscarPorID(inscripcionID)
	if !inscripcion.Existe() {
		return dto.NewAPIResponse(404, "Inscripción no encontrada", nil)
	}

	if inscripcion.EstaAprobada() {
		return dto.NewAPIResponse(200, "La inscripción ya ha sido aprobada.", nil)
	}

	exito := inscripcion.Aprobar()

	if !exito {
		return dto.NewAPIResponse(500, "Ha ocurrido un error interno en el sistema. Por favor, intenta nuevamente más tarde.", nil)
	}

	return dto.NewAPIResponse(200, "La inscripción ha sido aprobada exitosamente", nil)
}
