package usecase

import (
	"fmt"
	"lgc/src/domain"
	"lgc/src/view/dto"
	formrequest "lgc/src/view/form-request"
)

type InscripcionUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewInscripcionUseCase(inscripcionRepo domain.InscripcionRepository) *InscripcionUseCase {
	return &InscripcionUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *InscripcionUseCase) Execute(req formrequest.InscripcionFormRequest) dto.APIResponse {

	inscripcion := uc.inscripcionRepo.BuscarPorDocumento(req.Documento)
	if inscripcion.Existe() && inscripcion.EsValida() {
		return dto.NewAPIResponse(200, fmt.Sprintf("El documento %s ya cuenta con una inscripción registrada para el evento.", req.Documento), inscripcion.ToDTO())
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

	exito := inscripcion.Crear()

	if !exito {
		return dto.NewAPIResponse(500, "Ha ocurrido un error interno en el sistema. Por favor, intenta nuevamente más tarde.", nil)
	}

	return dto.NewAPIResponse(201, "La inscripción ha sido registrada exitosamente y el comprobante de pago ha sido recibido. Agradecemos su participación en el evento.", nil)
}
