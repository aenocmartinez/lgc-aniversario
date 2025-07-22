package dao

import (
	"lgc/src/view/dto"

	"gorm.io/gorm"
)

type EstadisticasDao struct {
	db *gorm.DB
}

func NewEstadisticasDao(db *gorm.DB) *EstadisticasDao {
	return &EstadisticasDao{db: db}
}

func (i *EstadisticasDao) ObtenerResumenEstadisticasEvento(cupoMax int) dto.EstadisticaEventoDTO {
	db := i.db

	// 1. Cupo presencial ocupado
	var totalPresenciales int64
	db.Table("participantes").
		Where("modalidad = ? AND dias_asistencia = ? AND inscripcion_id IN (SELECT id FROM inscripciones WHERE estado != ?)",
			"presencial", "sabado", "Rechazada").
		Count(&totalPresenciales)

	// 2. Total por modalidad
	var modalidadResults []struct {
		Modalidad string
		Total     int
	}
	db.Table("participantes").
		Select("modalidad, COUNT(*) as total").
		Group("modalidad").
		Scan(&modalidadResults)

	totalPorModalidad := map[string]int{}
	for _, r := range modalidadResults {
		totalPorModalidad[r.Modalidad] = r.Total
	}

	// 3. Total por días de asistencia
	var diasAsistenciaResults []struct {
		Dia   string
		Total int
	}
	db.Table("participantes").
		Select("dias_asistencia as dia, COUNT(*) as total").
		Group("dias_asistencia").
		Scan(&diasAsistenciaResults)

	totalPorDiaAsistencia := map[string]int{}
	for _, r := range diasAsistenciaResults {
		totalPorDiaAsistencia[r.Dia] = r.Total
	}

	// 4. Estado por forma de pago
	var estadoPagoResults []struct {
		FormaPago string
		Estado    string
		Total     int
	}
	db.Table("inscripciones").
		Select("forma_pago, estado, COUNT(*) as total").
		Group("forma_pago, estado").
		Scan(&estadoPagoResults)

	estadoPorFormaPago := map[string]map[string]int{}
	for _, r := range estadoPagoResults {
		if _, ok := estadoPorFormaPago[r.FormaPago]; !ok {
			estadoPorFormaPago[r.FormaPago] = map[string]int{}
		}
		estadoPorFormaPago[r.FormaPago][r.Estado] = r.Total
	}

	// 5. Inscripciones por día
	var inscripcionesPorDia []dto.InscripcionesDiaDTO
	db.Table("inscripciones").
		Select("DATE(created_at) as fecha, COUNT(*) as total").
		Group("DATE(created_at)").
		Order("fecha").
		Scan(&inscripcionesPorDia)

	// 6. Participantes sin iglesia
	var totalSinIglesia int64
	db.Table("participantes").
		Where("TRIM(iglesia) = 'No asiste a una iglesia'").
		Count(&totalSinIglesia)

	return dto.EstadisticaEventoDTO{
		CupoMaximoPresencial:    cupoMax,
		CupoUtilizadoPresencial: int(totalPresenciales),
		CupoRestantePresencial:  cupoMax - int(totalPresenciales),
		TotalPorModalidad:       totalPorModalidad,
		TotalPorDiaAsistencia:   totalPorDiaAsistencia,
		EstadoPorFormaPago:      estadoPorFormaPago,
		TotalSinIglesia:         int(totalSinIglesia),
		InscripcionesPorDia:     inscripcionesPorDia,
	}

}
