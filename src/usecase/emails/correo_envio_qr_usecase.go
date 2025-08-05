package usecase

import (
	"fmt"

	"lgc/src/domain"
	"lgc/src/infraestructure/email"

	"github.com/skip2/go-qrcode"
)

type CorreoEnvioQRUseCase struct {
	participanteRepo domain.ParticipanteRepository
	emailService     *email.EmailService
}

func NewCorreoEnvioQRUseCase(
	participanteRepo domain.ParticipanteRepository,
	emailService *email.EmailService,
) *CorreoEnvioQRUseCase {
	return &CorreoEnvioQRUseCase{
		participanteRepo: participanteRepo,
		emailService:     emailService,
	}
}

func (uc *CorreoEnvioQRUseCase) Execute() error {
	participantes := uc.participanteRepo.ObtenerParticipantesParaEnvioQR()

	for index, p := range participantes {

		// textoQR := fmt.Sprintf("https://dockerapps.pulzo.com/lgc-aniversario/participantes/buscar?documento=%s", p.GetDocumento())
		textoQR := fmt.Sprintf("https://dockerapps.pulzo.com/lgc-aniversario/participantes/visualizar?documento=%s", p.GetDocumento())

		qrBytes, err := qrcode.Encode(textoQR, qrcode.Medium, 256)
		if err != nil {
			return fmt.Errorf("error generando QR para %s: %w", p.GetDocumento(), err)
		}

		// HTML usa la imagen referenciada por CID
		qrHTML := `<img src="cid:qr-code.png" alt="Código QR" style="width:200px;height:200px;">`

		// htmlBody := fmt.Sprintf(`
		// 	<p>Hola %s,</p>
		// 	<p>Este es tu código QR para ingresar al evento:</p>
		// 	<p>%s</p>
		// 	<p>Bendiciones,<br>La Iglesia</p>
		// `, p.GetNombre(), qrHTML)
		htmlBody := fmt.Sprintf(`
			<p>Hola %s,</p>

			<p>Te compartimos tu <strong>código QR</strong>, el cual deberás presentar al ingresar al evento <strong>25° Aniversario – Llenos del Espíritu Santo</strong> de la Comunidad Cristiana Integral – La Gran Comisión.</p>

			<p>Este código es tu comprobante personal para validar tu asistencia en los puntos de ingreso.</p>

			<p>%s</p>

			<p>📅 <strong>Fechas del evento:</strong> 22, 23 y 24 de agosto de 2025</p>
			<p>❗ <strong>Nota:</strong> El ingreso el <strong>sábado 23</strong> de agosto es exclusivo para quienes completaron el proceso de pago.</p>

			<p>Si tienes alguna inquietud, contáctanos a <a href="mailto:grancomisionccieventos@gmail.com">grancomisionccieventos@gmail.com</a> o al WhatsApp 316 6972613.</p>

			<p>Con cariño,<br>Equipo de Eventos – La Gran Comisión</p>
		`, p.GetNombre(), qrHTML)

		// Llamamos a la nueva función que usa el Content-ID
		err = uc.emailService.EnviarEmailConQRUsandoCID("aenoc.martinez@gmail.com", "Código QR para ingreso al 25° Aniversario - La Gran Comisión", htmlBody, qrBytes)
		if err != nil {
			return fmt.Errorf("error enviando correo a %s: %w", p.GetEmail(), err)
		}

		if index == 9 {
			break
		}
	}

	return nil
}
