package usecase

import (
	"lgc/src/domain"
	"lgc/src/infraestructure/util"
	"lgc/src/view/dto"
	formrequest "lgc/src/view/form-request"
	"os"
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

	cupoMax, err := util.ConvertStringToInt(os.Getenv("APP_CUPO_MAX"))
	if err != nil {
		cupoMax = 400
	}

	inscripcion := domain.NewInscripcion(uc.inscripcionRepo)
	inscripcion.SetFormaPago(req.FormaPago)
	inscripcion.SetMontoPagoCOP(req.MontoCOP)
	inscripcion.SetMontoPagoUSD(req.MontoUSD)
	inscripcion.SetUrlSoportePago(req.UrlSoportePago)
	inscripcion.SetEstado("PreAprobada")

	if req.FormaPago == "efectivo" {
		inscripcion.SetEstado("Aprobada")
	}

	if req.FormaPago != "gratuito" && req.UrlSoportePago == "" {
		return dto.NewAPIResponse(400, "Se requiere el soporte de pago cuando la forma de pago no es gratuita.", nil)
	}

	var participantes []domain.Participante
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
		participantes = append(participantes, *participante)
	}

	err = uc.inscripcionRepo.CrearConValidacionDeCupo(inscripcion, participantes, cupoMax)
	if err != nil {
		if err.Error() == "cupo lleno" {
			return dto.NewAPIResponse(200, "El cupo máximo presencial ha sido alcanzado. No es posible registrar más participantes presenciales.", nil)
		}
		return dto.NewAPIResponse(500, "Ocurrió un error al registrar la inscripción. Intente nuevamente.", nil)
	}

	return dto.NewAPIResponse(201, "La inscripción ha sido registrada exitosamente. Agradecemos su participación en el evento.", nil)
}
