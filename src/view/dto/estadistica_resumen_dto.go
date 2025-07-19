package dto

type InscripcionesDiaDTO struct {
	Fecha string `json:"fecha"`
	Total int    `json:"total"`
}

type EstadisticaResumenDTO struct {
	TotalPorEstado               map[string]int            `json:"total_por_estado"`
	TotalPorAsistencia           map[string]int            `json:"total_por_asistencia"`
	TotalSinIglesia              int                       `json:"total_sin_iglesia"`
	DistribucionEstadoAsistencia map[string]map[string]int `json:"distribucion_estado_asistencia"`
	InscripcionesPorDia          []InscripcionesDiaDTO     `json:"inscripciones_por_dia"`
	CupoRestantePresencial       int                       `json:"cupo_restante_presencial"`
}
