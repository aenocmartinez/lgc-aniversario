package dao

import (
	"fmt"
	"lgc/src/domain"
	"lgc/src/view/dto"
	"log"
	"time"

	"gorm.io/gorm"
)

type InscripcionDao struct {
	db *gorm.DB
}

func NewInscripcionDao(db *gorm.DB) *InscripcionDao {
	return &InscripcionDao{db: db}
}

type inscripcionModel struct {
	ID             int64     `gorm:"primaryKey;column:id"`
	FormaPago      string    `gorm:"column:forma_pago"`
	MontoPagoCOP   int       `gorm:"column:monto_pagado_cop"`
	MontoPagoUSD   int       `gorm:"column:monto_pagado_usd"`
	UrlSoportePago string    `gorm:"column:soporte_pago_url"`
	Estado         string    `gorm:"column:estado"`
	FechaCreacion  time.Time `gorm:"column:created_at;<-:false"`
}

func (inscripcionModel) TableName() string {
	return "inscripciones"
}

func (i *InscripcionDao) Crear(inscripcion *domain.Inscripcion) bool {
	model := inscripcionModel{
		FormaPago:      inscripcion.GetFormaPago(),
		MontoPagoCOP:   inscripcion.GetMontoPagoCOP(),
		MontoPagoUSD:   inscripcion.GetMontoPagoUSD(),
		UrlSoportePago: inscripcion.GetUrlSoportePago(),
		Estado:         "PreAprobada",
	}

	result := i.db.Create(&model)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	}

	inscripcion.SetID(model.ID)
	inscripcion.SetEstado(model.Estado)

	return true
}

func (dao *InscripcionDao) CrearConValidacionDeCupo(
	inscripcion *domain.Inscripcion,
	participantes []domain.Participante,
	cupoMax int,
) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		// ✅ 1. Contar ocupados: solo presencial del sábado y que la inscripción no esté rechazada
		var ocupados int64
		tx.Raw(`
			SELECT COUNT(*)
			FROM participantes p
			INNER JOIN inscripciones i ON i.id = p.inscripcion_id
			WHERE p.modalidad = 'presencial' AND p.dias_asistencia = 'sabado' AND i.estado != 'Rechazada'
			FOR UPDATE
		`).Scan(&ocupados)

		// ✅ 2. Contar los nuevos participantes que aplican a esta condición
		var solicitados int
		for _, p := range participantes {
			if p.GetModalidad() == "presencial" && p.GetDiasAsistencia() == "sabado" {
				solicitados++
			}
		}

		// ✅ 3. Validar cupo
		if int(ocupados)+solicitados > cupoMax {
			return fmt.Errorf("cupo lleno")
		}

		// ✅ 4. Crear inscripción
		model := inscripcionModel{
			FormaPago:      inscripcion.GetFormaPago(),
			MontoPagoCOP:   inscripcion.GetMontoPagoCOP(),
			MontoPagoUSD:   inscripcion.GetMontoPagoUSD(),
			UrlSoportePago: inscripcion.GetUrlSoportePago(),
			Estado:         inscripcion.GetEstado(),
		}
		if err := tx.Create(&model).Error; err != nil {
			return err
		}
		inscripcion.SetID(model.ID)

		// ✅ 5. Crear participantes
		for _, p := range participantes {
			if err := tx.Exec(`
				INSERT INTO participantes (
					inscripcion_id, nombre_completo, numero_documento, correo_electronico,
					telefono, modalidad, dias_asistencia, iglesia, ciudad, autorizacion_datos
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				model.ID,
				p.GetNombre(),
				p.GetDocumento(),
				p.GetEmail(),
				p.GetTelefono(),
				p.GetModalidad(),
				p.GetDiasAsistencia(),
				p.GetIglesia(),
				p.GetCiudad(),
				p.GetHabeasData(),
			).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (i *InscripcionDao) Listar() []domain.Inscripcion {
	var models []inscripcionModel
	var resultado []domain.Inscripcion

	i.db.Order("id desc").Find(&models)

	for _, m := range models {
		ins := domain.NewInscripcion(NewInscripcionDao(i.db))
		ins.SetID(m.ID)
		ins.SetFormaPago(m.FormaPago)
		ins.SetMontoPagoCOP(m.MontoPagoCOP)
		ins.SetMontoPagoUSD(m.MontoPagoUSD)
		ins.SetUrlSoportePago(m.UrlSoportePago)
		ins.SetEstado(m.Estado)
		ins.SetFechaCreacion(m.FechaCreacion.Format("2006-01-02 15:04:05"))
		resultado = append(resultado, *ins)
	}

	return resultado
}

func (i *InscripcionDao) ListarConParticipantes() []dto.InscripcionConParticipantesDTO {
	var inscripciones []inscripcionModel
	var resultado []dto.InscripcionConParticipantesDTO

	// Obtener todas las inscripciones ordenadas por ID descendente
	i.db.Order("id desc").Find(&inscripciones)

	for _, ins := range inscripciones {
		// Obtener los participantes asociados a esta inscripción
		rows, err := i.db.Raw(`
			SELECT nombre_completo, numero_documento, correo_electronico,
			       telefono, modalidad, dias_asistencia, iglesia, ciudad, autorizacion_datos
			FROM participantes
			WHERE inscripcion_id = ?`, ins.ID).Rows()
		if err != nil {
			continue
		}

		defer rows.Close()

		var participantes []dto.ParticipanteDTO

		for rows.Next() {
			var p dto.ParticipanteDTO
			rows.Scan(&p.Nombre, &p.Documento, &p.Email, &p.Telefono, &p.Modalidad, &p.DiasAsistencia, &p.Iglesia, &p.Ciudad, &p.HabeasData)
			participantes = append(participantes, p)
		}

		resultado = append(resultado, dto.InscripcionConParticipantesDTO{
			ID:             ins.ID,
			FormaPago:      ins.FormaPago,
			MontoPagoCOP:   ins.MontoPagoCOP,
			MontoPagoUSD:   ins.MontoPagoUSD,
			UrlSoportePago: ins.UrlSoportePago,
			Estado:         ins.Estado,
			FechaCreacion:  ins.FechaCreacion.Format("2006-01-02 15:04:05"),
			Participantes:  participantes,
		})
	}

	return resultado
}

func (i *InscripcionDao) BuscarPorID(inscripcionID int64) domain.Inscripcion {
	var model inscripcionModel

	result := i.db.First(&model, "id = ?", inscripcionID)
	if result.Error != nil || result.RowsAffected == 0 {
		return domain.Inscripcion{}
	}

	ins := domain.NewInscripcion(i)
	ins.SetID(model.ID)
	ins.SetFormaPago(model.FormaPago)
	ins.SetMontoPagoCOP(model.MontoPagoCOP)
	ins.SetMontoPagoUSD(model.MontoPagoUSD)
	ins.SetUrlSoportePago(model.UrlSoportePago)
	ins.SetEstado(model.Estado)
	ins.SetFechaCreacion(model.FechaCreacion.Format("2006-01-02 15:04:05"))

	return *ins
}

func (i *InscripcionDao) AgregarParticipante(inscripcionID int64, participante domain.Participante) bool {
	result := i.db.Exec(`
		INSERT INTO participantes (
			inscripcion_id, nombre_completo, numero_documento, correo_electronico,
			telefono, modalidad, dias_asistencia, iglesia, ciudad, autorizacion_datos
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		inscripcionID,
		participante.GetNombre(),
		participante.GetDocumento(),
		participante.GetEmail(),
		participante.GetTelefono(),
		participante.GetModalidad(),
		participante.GetDiasAsistencia(),
		participante.GetIglesia(),
		participante.GetCiudad(),
		participante.GetHabeasData(),
	)

	log.Println(result.Error)

	return result.Error == nil && result.RowsAffected > 0
}

func (i *InscripcionDao) ObtenerParticipantes(inscripcionID int64) []domain.Participante {
	var participantes []domain.Participante

	rows, err := i.db.Raw(`
		SELECT nombre_completo, numero_documento, correo_electronico,
		       telefono, modalidad, dias_asistencia, iglesia, ciudad, autorizacion_datos
		FROM participantes
		WHERE inscripcion_id = ?`, inscripcionID).Rows()
	if err != nil {
		return participantes
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.Participante
		var nombre, doc, email, tel, modalidad, dias, iglesia, ciudad string
		var habeas bool

		rows.Scan(&nombre, &doc, &email, &tel, &modalidad, &dias, &iglesia, &ciudad, &habeas)

		p.SetNombre(nombre)
		p.SetDocumento(doc)
		p.SetEmail(email)
		p.SetTelefono(tel)
		p.SetModalidad(modalidad)
		p.SetDiasAsistencia(dias)
		p.SetIglesia(iglesia)
		p.SetCiudad(ciudad)
		p.SetHabeasData(habeas)

		participantes = append(participantes, p)
	}

	return participantes
}

func (i *InscripcionDao) Aprobar(inscripcionID int64) bool {
	result := i.db.Model(&inscripcionModel{}).
		Where("id = ?", inscripcionID).
		Update("estado", "Aprobada")

	return result.Error == nil && result.RowsAffected > 0
}

func (i *InscripcionDao) Rechazar(inscripcionID int64) bool {
	result := i.db.Model(&inscripcionModel{}).
		Where("id = ?", inscripcionID).
		Update("estado", "Rechazada")

	return result.Error == nil && result.RowsAffected > 0
}

func (dao *InscripcionDao) CuposDisponibles(cupoMax int) (int, int) {
	var ocupados int64

	dao.db.Raw(`
		SELECT COUNT(*) FROM participantes p
		INNER JOIN inscripciones i ON i.id = p.inscripcion_id
		WHERE p.modalidad = 'presencial' AND i.estado != 'Rechazada'
	`).Scan(&ocupados)

	return int(ocupados), cupoMax - int(ocupados)
}

func (dao *InscripcionDao) BuscarParticipantePorDocumento(documento string) (*domain.Participante, string, error) {
	var result struct {
		ID            int64
		InscripcionID int64
		Estado        string
	}

	err := dao.db.Raw(`
		SELECT p.id, p.inscripcion_id, i.estado
		FROM participantes p
		INNER JOIN inscripciones i ON i.id = p.inscripcion_id
		WHERE p.numero_documento = ?
		LIMIT 1
	`, documento).Scan(&result).Error

	if err != nil {
		return nil, "", err
	}
	if result.ID == 0 {
		return nil, "", nil // No encontrado
	}

	participante := domain.NewParticipante(nil)
	participante.SetID(result.ID)

	inscripcion := dao.BuscarPorID(result.InscripcionID)

	participante.SetInscripcion(&inscripcion)
	// participante.SetInscripcionID(result.InscripcionID)

	return participante, result.Estado, nil
}

func (dao *InscripcionDao) EliminarParticipanteYValidarInscripcion(participanteID int64) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		var inscripcionID int64
		err := tx.Raw(`SELECT inscripcion_id FROM participantes WHERE id = ?`, participanteID).Scan(&inscripcionID).Error
		if err != nil {
			log.Println("SELECT inscripción_id error:", err)
			return err
		}
		if inscripcionID == 0 {
			log.Println("No se encontró inscripción asociada al participante")
			return fmt.Errorf("inscripción no encontrada para el participante")
		}
		log.Println("inscripcionID:", inscripcionID)

		var total int64
		err = tx.Raw(`SELECT COUNT(*) FROM participantes WHERE inscripcion_id = ?`, inscripcionID).Scan(&total).Error
		if err != nil {
			log.Println("SELECT COUNT error:", err)
			return err
		}
		log.Println("Total participantes en inscripción:", total)

		if err := tx.Exec(`DELETE FROM participantes WHERE id = ?`, participanteID).Error; err != nil {
			log.Println("Error al eliminar participante:", err)
			return err
		}
		log.Println("Participante eliminado con éxito")

		if total == 1 {
			if err := tx.Exec(`DELETE FROM inscripciones WHERE id = ?`, inscripcionID).Error; err != nil {
				log.Println("Error al eliminar inscripción:", err)
				return err
			}
			log.Println("Inscripción también eliminada")
		}

		return nil
	})
}
