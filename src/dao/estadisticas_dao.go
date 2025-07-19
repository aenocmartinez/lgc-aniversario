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

func (i *EstadisticasDao) ObtenerResumenEstadisticas(cupoMax int) dto.EstadisticaResumenDTO {
	db := i.db

	// 1. Total por estado
	var estadoResults []struct {
		Estado string
		Total  int
	}
	db.Model(&formularioDB{}).
		Select("estado, COUNT(*) as total").
		Group("estado").
		Scan(&estadoResults)

	totalPorEstado := map[string]int{}
	for _, r := range estadoResults {
		totalPorEstado[r.Estado] = r.Total
	}

	// 2. Total por asistencia
	var asistenciaResults []struct {
		Asistencia string
		Total      int
	}
	db.Model(&formularioDB{}).
		Select("asistencia, COUNT(*) as total").
		Group("asistencia").
		Scan(&asistenciaResults)

	totalPorAsistencia := map[string]int{}
	for _, r := range asistenciaResults {
		totalPorAsistencia[r.Asistencia] = r.Total
	}

	// 3. Total sin iglesia
	var totalSinIglesia int64
	db.Model(&formularioDB{}).Where("TRIM(iglesia) = 'No asiste a una iglesia'").Count(&totalSinIglesia)

	// 4. Distribución por estado y asistencia
	var estadoAsistenciaResults []struct {
		Estado     string
		Asistencia string
		Total      int
	}
	db.Model(&formularioDB{}).
		Select("estado, asistencia, COUNT(*) as total").
		Group("estado, asistencia").
		Scan(&estadoAsistenciaResults)

	distribucion := map[string]map[string]int{}
	for _, r := range estadoAsistenciaResults {
		if distribucion[r.Estado] == nil {
			distribucion[r.Estado] = map[string]int{}
		}
		distribucion[r.Estado][r.Asistencia] = r.Total
	}

	// 5. Inscripciones por día
	var inscripcionesPorDia []dto.InscripcionesDiaDTO
	db.Model(&formularioDB{}).
		Select("DATE(fecha_registro) as fecha, COUNT(*) as total").
		Group("DATE(fecha_registro)").
		Order("fecha").
		Scan(&inscripcionesPorDia)

	// 6. Cupo restante presencial
	var totalPresenciales int64
	db.Model(&formularioDB{}).
		Where("asistencia = 'Presencial' AND estado != 'Anulada'").
		Count(&totalPresenciales)

	return dto.EstadisticaResumenDTO{
		TotalPorEstado:               totalPorEstado,
		TotalPorAsistencia:           totalPorAsistencia,
		TotalSinIglesia:              int(totalSinIglesia),
		DistribucionEstadoAsistencia: distribucion,
		InscripcionesPorDia:          inscripcionesPorDia,
		CupoRestantePresencial:       cupoMax - int(totalPresenciales),
	}
}
