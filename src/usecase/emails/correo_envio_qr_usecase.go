package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"lgc/src/domain"
	"lgc/src/infraestructure/email"

	"github.com/skip2/go-qrcode"
)

type repoMarcaEnvio interface {
	MarcarEstadoEnvioQR(documento string, exito bool, detalle string, enviadoAt *time.Time) error
}

type CorreoEnvioQRUseCase struct {
	participanteRepo domain.ParticipanteRepository
	emailService     *email.EmailService

	BatchSize int // éxitos por ejecución
	Workers   int // goroutines consumidoras
}

func NewCorreoEnvioQRUseCase(
	participanteRepo domain.ParticipanteRepository,
	emailService *email.EmailService,
) *CorreoEnvioQRUseCase {
	return &CorreoEnvioQRUseCase{
		participanteRepo: participanteRepo,
		emailService:     emailService,
		BatchSize:        50,
		Workers:          3,
	}
}

func (uc *CorreoEnvioQRUseCase) Execute() error {
	start := time.Now()

	// *** NO TOCAR ***
	participantes := uc.participanteRepo.ObtenerParticipantesParaEnvioQR()
	if len(participantes) == 0 {
		log.Println("[init] No hay participantes para enviar.")
		return nil
	}

	repoMarca, tieneMarca := uc.participanteRepo.(repoMarcaEnvio)

	maxExitos := uc.BatchSize
	if maxExitos <= 0 {
		maxExitos = 50
	}
	workers := uc.Workers
	if workers <= 0 {
		workers = 3
	}

	// Filtrar en memoria los ya enviados (no consumen cupo)
	elegibles := make([]domain.Participante, 0, len(participantes))
	for _, p := range participantes {
		if p.YaFueEnviadoQR() {
			continue
		}
		elegibles = append(elegibles, p)
	}

	total := len(participantes)
	totalElegibles := len(elegibles)
	if totalElegibles == 0 {
		log.Printf("[init] Total=%d, elegibles=%d → nada que enviar.\n", total, totalElegibles)
		return nil
	}
	log.Printf("[init] Total=%d, elegibles=%d, tope=%d, workers=%d\n", total, totalElegibles, maxExitos, workers)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan domain.Participante, workers*2)
	var wg sync.WaitGroup
	var exitos int32
	var errores int32

	// Progreso periódico
	doneProgress := make(chan struct{})
	go func() {
		t := time.NewTicker(2 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				ok := atomic.LoadInt32(&exitos)
				errs := atomic.LoadInt32(&errores)
				restantes := maxExitos - int(ok)
				if restantes < 0 {
					restantes = 0
				}
				log.Printf("[progress] éxitos=%d/%d, errores=%d, restantes=%d\n", ok, maxExitos, errs, restantes)
			case <-doneProgress:
				return
			}
		}
	}()

	// Workers
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for p := range jobs {
				// Corte temprano si ya alcanzamos tope
				if atomic.LoadInt32(&exitos) >= int32(maxExitos) {
					log.Printf("[worker %d] tope alcanzado, saliendo.\n", id)
					return
				}

				if err := uc.enviarUno(ctx, p, repoMarca, tieneMarca, id); err == nil {
					// Éxito → cuenta y quizá cancela
					if atomic.AddInt32(&exitos, 1) >= int32(maxExitos) {
						log.Printf("[worker %d] tope alcanzado tras este éxito, cancelando.\n", id)
						cancel()
						return
					}
				} else {
					atomic.AddInt32(&errores, 1)
				}
			}
			log.Printf("[worker %d] canal cerrado, fin.\n", id)
		}(w + 1)
	}

	// FEED LOOP: select sin default (no se bloquea si se cancela)
feedLoop:
	for _, p := range elegibles {
		select {
		case <-ctx.Done():
			log.Println("[feed] cancelado por tope; dejo de encolar trabajos.")
			break feedLoop
		case jobs <- p:
			// encolado OK
		}
	}
	close(jobs)
	wg.Wait()
	close(doneProgress)

	dur := time.Since(start)
	ok := atomic.LoadInt32(&exitos)
	errs := atomic.LoadInt32(&errores)
	restantes := maxExitos - int(ok)
	if restantes < 0 {
		restantes = 0
	}
	log.Printf("[done] éxitos=%d/%d, errores=%d, restantes=%d, elegibles=%d, dur=%s\n",
		ok, maxExitos, errs, restantes, totalElegibles, dur)

	return nil
}

func (uc *CorreoEnvioQRUseCase) enviarUno(
	ctx context.Context,
	p domain.Participante,
	repoMarca repoMarcaEnvio,
	tieneMarca bool,
	workerID int,
) error {
	to := p.GetEmail()
	if to == "" {
		log.Printf("[warn] worker=%d doc=%s sin email, se omite.\n", workerID, p.GetDocumento())
		return errors.New("participante sin email")
	}

	select {
	case <-ctx.Done():
		return context.Canceled
	default:
	}

	// Construcción del QR
	textoQR := fmt.Sprintf(
		"https://dockerapps.pulzo.com/lgc-aniversario/participantes/visualizar?documento=%s",
		p.GetDocumento(),
	)
	qrBytes, err := qrcode.Encode(textoQR, qrcode.Medium, 256)
	if err != nil {
		log.Printf("[error] worker=%d doc=%s generando QR: %v\n", workerID, p.GetDocumento(), err)
		return fmt.Errorf("qr: %w", err)
	}

	qrHTML := `<img src="cid:qr-code.png" alt="Código QR" style="width:200px;height:200px;">`
	htmlBody := fmt.Sprintf(`
		<p>Hola %s,</p>

		<p>Te compartimos tu <strong>código QR</strong>, el cual deberás presentar al ingresar al evento <strong>25° Aniversario – Llenos del Espíritu Santo</strong> de la Comunidad Cristiana Integral – La Gran Comisión.</p>

		<p><strong>Este código es tu comprobante personal para validar tu asistencia en los puntos de ingreso.</strong></p>

		<p>%s</p>

		<p>📅 <strong>Fechas del evento:</strong> 22, 23 y 24 de agosto de 2025</p>
		<p>❗ <strong>Nota:</strong> El ingreso el <strong>sábado 23</strong> de agosto es exclusivo para quienes completaron el proceso de pago.</p>

		<p>Si tienes alguna inquietud, contáctanos a 
		<a href="mailto:grancomisionccieventos@gmail.com">grancomisionccieventos@gmail.com</a> 
		o al WhatsApp <strong>316 6972613</strong>.</p>

		<p><br>Comunidad Cristiana Integral – La Gran Comisión</p>
	`, p.GetNombre(), qrHTML)

	// Reintento simple (2 intentos)
	const maxIntentos = 2
	for intento := 1; intento <= maxIntentos; intento++ {
		select {
		case <-ctx.Done():
			log.Printf("[cancel] worker=%d doc=%s cancelado antes de enviar.\n", workerID, p.GetDocumento())
			return context.Canceled
		default:
		}

		log.Printf("[send] worker=%d doc=%s to=%s intento=%d/%d\n", workerID, p.GetDocumento(), to, intento, maxIntentos)

		err = uc.emailService.EnviarEmailConQRUsandoCID(
			to,
			"Código QR para ingreso al 25° Aniversario - La Gran Comisión",
			htmlBody,
			qrBytes,
		)
		if err == nil {
			// Marca solo en éxito (flag + fecha/hora)
			if tieneMarca && repoMarca != nil {
				loc, lerr := time.LoadLocation("America/Bogota")
				var nowPtr *time.Time
				if lerr == nil {
					now := time.Now().In(loc)
					nowPtr = &now
				}
				if mErr := repoMarca.MarcarEstadoEnvioQR(p.GetDocumento(), true, "OK", nowPtr); mErr != nil {
					log.Printf("[warn] worker=%d doc=%s enviado OK, pero falló marcar estado: %v\n", workerID, p.GetDocumento(), mErr)
				}
			}
			log.Printf("[ok] worker=%d doc=%s enviado.\n", workerID, p.GetDocumento())
			return nil
		}

		log.Printf("[retry] worker=%d doc=%s intento=%d err=%v\n", workerID, p.GetDocumento(), intento, err)
		time.Sleep(2 * time.Second)
	}

	log.Printf("[fail] worker=%d doc=%s después de %d intentos. Último error: %v\n", workerID, p.GetDocumento(), maxIntentos, err)
	return err
}
