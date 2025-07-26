package usecase

import (
	"lgc/src/domain"
	"lgc/src/infraestructure/email"
	"lgc/src/infraestructure/util"
	"log"
)

type SendEmailUseCase struct {
	emailService *email.EmailService
}

func NewSendEmailUseCase(emailService *email.EmailService) *SendEmailUseCase {
	return &SendEmailUseCase{emailService: emailService}
}

func (uc *SendEmailUseCase) Execute(inscripcion *domain.Inscripcion) error {

	var asunto string = "Confirmación de inscripción – Aniversario #25 Iglesia La Gran Comisión"
	var archivoAdjunto string = "resources/Programacion_Aniversario_25_La_Gran_Comision.pdf"

	for _, participante := range inscripcion.Participantes() {

		cuerpoCorreo := GetCorreoInscripcionRealizada(util.ToCapitalCase(participante.GetNombre()),
			participante.GetModalidad(),
			participante.GetDiasAsistencia())

		// err := uc.emailService.EnviarEmail(participante.GetEmail(), asunto, cuerpoCorreo)
		err := uc.emailService.EnviarEmailConAdjunto(participante.GetEmail(), asunto, cuerpoCorreo, archivoAdjunto)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}
