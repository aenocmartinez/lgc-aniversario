package dao

import (
	"lgc/src/domain"
	"lgc/src/view/dto"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ParticipanteDao struct {
	db *gorm.DB
}

func NewParticipanteDao(db *gorm.DB) *ParticipanteDao {
	return &ParticipanteDao{db: db}
}

func (r *ParticipanteDao) ObtenerParticipantesParaLogistica() []dto.ReporteLogisticaDTO {
	var results []dto.ReporteLogisticaDTO

	r.db.Table("inscripciones AS i").
		Select("p.nombre_completo, p.numero_documento, p.correo_electronico, p.dias_asistencia, p.telefono, p.iglesia").
		Joins("INNER JOIN participantes p ON i.id = p.inscripcion_id").
		Where("i.estado <> ?", "Rechazada").
		Where("p.modalidad <> ?", "Virtual").
		Scan(&results)

	for i := range results {
		switch strings.ToLower(strings.TrimSpace(results[i].DiasAsistencia)) {
		case "sabado":
			results[i].DiasAsistencia = "viernes, sábado y domingo"
		case "viernes_y_domingo":
			results[i].DiasAsistencia = "viernes y domingo"
		}
	}

	return results
}

func (r *ParticipanteDao) ObtenerParticipantesParaEnvioQR() []domain.Participante {
	type rowT struct {
		ID             int64
		Nombre         string
		Documento      string
		Email          string
		DiasAsistencia string
		QREnviado      bool
		QREnviadoAt    *time.Time
	}

	var rows []rowT

	r.db.Table("participantes AS p").
		Select(`
			p.id                                   AS id,
			p.nombre_completo                      AS nombre,
			p.numero_documento                     AS documento,
			p.correo_electronico                   AS email,
			p.dias_asistencia                      AS dias_asistencia,
			p.qr_enviado                           AS qr_enviado,
			p.qr_enviado_at                        AS qr_enviado_at
		`).
		Joins("INNER JOIN inscripciones i ON i.id = p.inscripcion_id").
		Where("i.estado = ?", "Aprobada").
		Where("LOWER(p.modalidad) <> ?", "virtual").
		Where("p.qr_enviado = ?", false).
		Order("p.id ASC").
		Scan(&rows)

	participantes := make([]domain.Participante, 0, len(rows))
	for _, rrow := range rows {
		p := domain.NewParticipante(r)
		p.SetID(rrow.ID)
		p.SetNombre(rrow.Nombre)
		p.SetDocumento(rrow.Documento)
		p.SetEmail(strings.TrimSpace(rrow.Email))
		p.SetDiasAsistencia(strings.TrimSpace(rrow.DiasAsistencia))

		p.SetQREnviado(rrow.QREnviado)
		p.SetQREnviadoAt(rrow.QREnviadoAt)

		participantes = append(participantes, *p)
	}

	return participantes
}

func (r *ParticipanteDao) BuscarParticipantePorDocumento(documento string) domain.Participante {
	var row struct {
		ID                   int64
		Nombre               string
		Documento            string
		Modalidad            string
		DiasAsistencia       string
		QREnviado            bool
		QREnviadoAt          *time.Time
		AsistenciaConfirmada bool
	}

	r.db.Table("participantes AS p").
		Select(`
			p.id,
			p.nombre_completo AS nombre,
			p.numero_documento AS documento,
			p.modalidad,
			p.dias_asistencia,
			p.qr_enviado,
			p.qr_enviado_at,
			EXISTS (
				SELECT 1 FROM participante_asistencia pa
				WHERE pa.participante_id = p.id
			) AS asistencia_confirmada
		`).
		Joins("INNER JOIN inscripciones i ON i.id = p.inscripcion_id").
		Where("i.estado = ?", "Aprobada").
		Where("p.modalidad <> ?", "virtual").
		Where("p.numero_documento = ?", documento).
		Limit(1).
		Scan(&row)

	p := domain.NewParticipante(r)
	p.SetID(row.ID)
	p.SetNombre(row.Nombre)
	p.SetDocumento(row.Documento)
	p.SetModalidad(row.Modalidad)
	p.SetDiasAsistencia(row.DiasAsistencia)
	p.SetAsistenciaConfirmada(row.AsistenciaConfirmada)
	p.SetQREnviado(row.QREnviado)
	p.SetQREnviadoAt(row.QREnviadoAt)

	return *p
}

type participanteAsistencia struct {
	ID             int64 `gorm:"primaryKey"`
	ParticipanteID int64 `gorm:"column:participante_id;uniqueIndex:ux_participante_asistencia_participante"`
}

func (r *ParticipanteDao) RegistrarAsistencia(participanteID int64) bool {
	if participanteID <= 0 {
		return false
	}
	pa := participanteAsistencia{ParticipanteID: participanteID}

	err := r.db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "participante_id"}},
			DoNothing: true,
		},
	).Create(&pa).Error

	return err == nil
}

func (r *ParticipanteDao) MarcarEstadoEnvioQR(documento string, exito bool, detalle string, enviadoAt *time.Time) error {
	updates := map[string]any{
		"qr_enviado":         exito,
		"qr_enviado_detalle": detalle,
	}

	if exito {
		if enviadoAt != nil {
			updates["qr_enviado_at"] = *enviadoAt
		} else {
			updates["qr_enviado_at"] = gorm.Expr("CURRENT_TIMESTAMP")
		}

		if detalle == "" {
			updates["qr_enviado_detalle"] = "OK"
		}
	}

	tx := r.db.Table("participantes").
		Where("numero_documento = ?", documento).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Updates(updates)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *ParticipanteDao) GetAsistenciasConfirmadas() int64 {
	var count int64
	r.db.Table("participante_asistencia").Count(&count)
	return count
}
