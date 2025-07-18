package usecase

import (
	"fmt"
	"lgc/src/domain"
	"lgc/src/view/dto"
	formrequest "lgc/src/view/form-request"
)

type RealizarInscripcionUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewRealizarInscripcionUseCase(inscripcionRepo domain.InscripcionRepository) *RealizarInscripcionUseCase {
	return &RealizarInscripcionUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *RealizarInscripcionUseCase) Execute(req formrequest.InscripcionFormRequest) dto.APIResponse {

	inscripcion := uc.inscripcionRepo.BuscarPorDocumento(req.Documento)
	if inscripcion.Existe() {
		if inscripcion.EsValida() {
			return dto.NewAPIResponse(
				200,
				fmt.Sprintf("El documento %s ya tiene una inscripción validada. Su participación en el evento ha sido confirmada.", req.Documento),
				inscripcion.ToDTO(),
			)
		} else {
			return dto.NewAPIResponse(
				200,
				fmt.Sprintf("El documento %s ya cuenta con una inscripción registrada y se encuentra en proceso de validación.", req.Documento),
				inscripcion.ToDTO(),
			)
		}

	}

	inscripcion.SetNombre(req.Nombre)
	inscripcion.SetDocumento(req.Documento)
	inscripcion.SetEmail(req.Email)
	inscripcion.SetTelefono(req.Telefono)
	inscripcion.SetHabeasData(req.HabeasData)
	inscripcion.SetComprobantePago(req.ComprobatePago)
	inscripcion.SetAsistencia(req.Asistencia)
	inscripcion.SetCiudad(req.Ciudad)
	inscripcion.SetIglesia(req.Iglesia)
	inscripcion.SetEstado("Pendiente")

	exito := inscripcion.Crear()

	if !exito {
		return dto.NewAPIResponse(500, "Ha ocurrido un error interno en el sistema. Por favor, intenta nuevamente más tarde.", nil)
	}

	return dto.NewAPIResponse(201, "La inscripción ha sido registrada exitosamente y el comprobante de pago ha sido recibido. Agradecemos su participación en el evento.", nil)
}
