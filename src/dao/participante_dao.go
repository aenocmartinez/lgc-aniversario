package dao

import (
	"lgc/src/view/dto"
	"strings"

	"gorm.io/gorm"
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
		Select("p.nombre_completo, p.numero_documento, p.correo_electronico, p.dias_asistencia, p.telefono").
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
