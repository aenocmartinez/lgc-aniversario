package usecase

import (
	"fmt"
	"lgc/src/domain"
	"lgc/src/infraestructure/util"
	"lgc/src/view/dto"
	formrequest "lgc/src/view/form-request"
	"os"
)

type RealizarInscripcionPagoEnEfectivoUseCase struct {
	inscripcionRepo domain.InscripcionRepository
}

func NewRealizarInscripcionPagoEnEfectivoUseCase(inscripcionRepo domain.InscripcionRepository) *RealizarInscripcionPagoEnEfectivoUseCase {
	return &RealizarInscripcionPagoEnEfectivoUseCase{
		inscripcionRepo: inscripcionRepo,
	}
}

func (uc *RealizarInscripcionPagoEnEfectivoUseCase) Execute(req formrequest.InscripcionPagoEfectivoFormRequest) dto.APIResponse {
	inscripcion := uc.inscripcionRepo.BuscarPorDocumento(req.Documento)
	if inscripcion.Existe() {
		if inscripcion.EstaAprobada() {
			return dto.NewAPIResponse(
				200,
				fmt.Sprintf("El documento %s ya tiene una inscripción aprobada. Su participación en el evento ha sido confirmada.", req.Documento),
				inscripcion.ToDTO(),
			)
		}

		if inscripcion.EstaPreAprobada() {
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
	inscripcion.SetComprobantePago("Pago en efectivo")
	inscripcion.SetAsistencia("Presencial")
	inscripcion.SetCiudad("")
	inscripcion.SetIglesia(req.Iglesia)
	inscripcion.SetEstado("Aprobada")
	inscripcion.SetMedioPago("Efectivo")

	cupoMax, err := util.ConvertStringToInt(os.Getenv("APP_CUPO_MAX"))
	if err != nil {
		cupoMax = 400
	}

	ok, err := uc.inscripcionRepo.CrearConValidacionDeCupo(inscripcion, cupoMax)
	if err != nil {
		if err.Error() == "cupo lleno" {
			return dto.NewAPIResponse(200, "El cupo disponible ha sido completado, por lo tanto, no es posible procesar nuevas inscripciones.", nil)
		}
		return dto.NewAPIResponse(500, "Ha ocurrido un error al registrar la inscripción. Por favor, inténtelo de nuevo.", nil)
	}

	if ok {
		return dto.NewAPIResponse(201, "La inscripción ha sido registrada exitosamente. Agradecemos su participación en el evento.", nil)
	}

	return dto.NewAPIResponse(500, "No fue posible completar la inscripción. Inténtelo más tarde.", nil)
}
