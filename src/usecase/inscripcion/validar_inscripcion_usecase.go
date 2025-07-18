package usecase

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
)

type ValidarInscripcionUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewValidarInscripcionUseCase(inscripcionRepo domain.InscripcionRepository) *ValidarInscripcionUseCase {
	return &ValidarInscripcionUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *ValidarInscripcionUseCase) Execute(inscripcionID int64) dto.APIResponse {

	inscripcion := uc.inscripcionRepo.BuscarPorID(inscripcionID)
	if !inscripcion.Existe() {
		return dto.NewAPIResponse(404, "Inscripción no encontrada", nil)
	}

	if inscripcion.EsValida() {
		return dto.NewAPIResponse(200, "La inscripción ya ha sido validada.", nil)
	}

	exito := inscripcion.Aprobar()

	if !exito {
		return dto.NewAPIResponse(500, "Ha ocurrido un error interno en el sistema. Por favor, intenta nuevamente más tarde.", nil)
	}

	return dto.NewAPIResponse(200, "Inscripción ha sido validada exitosamente", nil)
}
