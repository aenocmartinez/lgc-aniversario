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

		textoQR := fmt.Sprintf("Nombre: %s\nDocumento: %s", p.GetNombre(), p.GetDocumento())

		qrBytes, err := qrcode.Encode(textoQR, qrcode.Medium, 256)
		if err != nil {
			return fmt.Errorf("error generando QR para %s: %w", p.GetDocumento(), err)
		}

		// HTML usa la imagen referenciada por CID
		qrHTML := `<img src="cid:qr-code.png" alt="Código QR" style="width:200px;height:200px;">`

		htmlBody := fmt.Sprintf(`
			<p>Hola %s,</p>
			<p>Este es tu código QR para ingresar al evento:</p>
			<p>%s</p>
			<p>Bendiciones,<br>La Iglesia</p>
		`, p.GetNombre(), qrHTML)

		// Llamamos a la nueva función que usa el Content-ID
		err = uc.emailService.EnviarEmailConQRUsandoCID("aenoc.martinez@gmail.com", "Tu QR para el evento", htmlBody, qrBytes)
		if err != nil {
			return fmt.Errorf("error enviando correo a %s: %w", p.GetEmail(), err)
		}

		if index == 0 {
			break
		}
	}

	return nil
}
