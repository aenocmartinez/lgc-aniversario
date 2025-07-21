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
	inscripcion := domain.NewInscripcion(uc.inscripcionRepo)
	inscripcion.SetFormaPago(req.FormaPago)
	inscripcion.SetMontoPagoCOP(req.MontoCOP)
	inscripcion.SetMontoPagoUSD(req.MontoUSD)
	inscripcion.SetUrlSoportePago(req.UrlSoportePago)

	inscripcion.SetEstado("PreAprobada")
	if req.FormaPago == "efectivo" {
		inscripcion.SetEstado("Aprobada")
	}

	if req.FormaPago != "gratuito" {
		if req.UrlSoportePago == "" {
			return dto.NewAPIResponse(400, "Se requiere el soporte de pago cuando la forma de pago no es gratuita.", nil)
		}
	}

	if !inscripcion.Crear() {
		return dto.NewAPIResponse(500, "No fue posible registrar la inscripción. Intente nuevamente.", nil)
	}

	for _, p := range req.Participantes {
		participante := domain.NewParticipante(nil)
		participante.SetNombre(p.Nombre)
		participante.SetDocumento(p.Documento)
		participante.SetEmail(p.Email)
		participante.SetTelefono(p.Telefono)
		participante.SetModalidad(p.Modalidad)
		participante.SetDiasAsistencia(p.DiasAsistencia)
		participante.SetIglesia(p.Iglesia)
		participante.SetCiudad(p.Ciudad)
		participante.SetHabeasData(p.HabeasData)

		if !inscripcion.AgregarParticipante(*participante) {
			return dto.NewAPIResponse(500, fmt.Sprintf("No fue posible registrar al participante %s", p.Nombre), nil)
		}
	}

	return dto.NewAPIResponse(201, "La inscripción ha sido registrada exitosamente. Agradecemos su participación en el evento.", nil)
}
