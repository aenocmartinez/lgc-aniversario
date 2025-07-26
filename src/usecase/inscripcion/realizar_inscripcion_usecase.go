package usecase

import (
	"fmt"
	"lgc/src/domain"
	"lgc/src/infraestructure/email"
	"lgc/src/infraestructure/util"
	usecase "lgc/src/usecase/emails"
	"lgc/src/view/dto"
	formrequest "lgc/src/view/form-request"
	"log"
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

	var participantes []domain.Participante
	for _, p := range req.Participantes {
		existente, estadoInscripcion, err := uc.inscripcionRepo.BuscarParticipantePorDocumento(p.Documento)
		if err != nil {
			return dto.NewAPIResponse(500, fmt.Sprintf("Error al verificar el documento %s", p.Documento), nil)
		}

		if existente != nil {
			formaPago := existente.GetInscripcion().GetFormaPago()

			log.Println("formaPago: ", formaPago)
			log.Println("estadoInscripcion: ", estadoInscripcion)

			if estadoInscripcion == "Aprobada" && formaPago != "gratuito" {
				msg := fmt.Sprintf("El participante con documento %s ya está aprobado y no puede volver a inscribirse.", p.Documento)
				return dto.NewAPIResponse(400, msg, nil)
			}

			if formaPago == "gratuito" {
				fmt.Println("Etrana Eliminar Participante")
				// Eliminar al participante para permitir nueva inscripción
				err := uc.inscripcionRepo.EliminarParticipanteYValidarInscripcion(existente.GetID())
				if err != nil {
					return dto.NewAPIResponse(500, fmt.Sprintf("Error al eliminar el participante anterior con documento %s", p.Documento), nil)
				}
			} else {
				msg := fmt.Sprintf("El participante con documento %s ya tiene una inscripción activa.", p.Documento)
				return dto.NewAPIResponse(400, msg, nil)
			}
		}

		// Crear nuevo participante
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

	// Crear inscripción
	inscripcion := domain.NewInscripcion(uc.inscripcionRepo)
	inscripcion.SetFormaPago(req.FormaPago)
	inscripcion.SetMontoPagoCOP(req.MontoCOP)
	inscripcion.SetMontoPagoUSD(req.MontoUSD)

	if req.FormaPago == "gratuito" {
		inscripcion.SetEstado("Aprobada")
	} else {
		inscripcion.SetEstado("PreAprobada")
	}

	if req.FormaPago != "efectivo" {
		inscripcion.SetUrlSoportePago(req.UrlSoportePago)
	} else {
		inscripcion.SetUrlSoportePago("Pago en efectivo")
	}

	err = uc.inscripcionRepo.CrearConValidacionDeCupo(inscripcion, participantes, cupoMax)
	if err != nil {
		if err.Error() == "cupo lleno" {
			return dto.NewAPIResponse(200, "El cupo máximo presencial ha sido alcanzado. No es posible registrar más participantes presenciales.", nil)
		}
		return dto.NewAPIResponse(500, "Ocurrió un error al registrar la inscripción. Intente nuevamente.", nil)
	}

	// Enviando correo
	sendEmail := usecase.NewSendEmailUseCase(email.NewEmailService(email.GetEmailConfig()))
	go sendEmail.Execute(inscripcion)

	return dto.NewAPIResponse(201, "La inscripción ha sido registrada exitosamente. Agradecemos su participación en el evento.", nil)
}
