package dto

type InscripcionesDiaDTO struct {
	Fecha string `json:"fecha"`
	Total int    `json:"total"`
}

type EstadisticaEventoDTO struct {
	CupoMaximoPresencial    int
	CupoUtilizadoPresencial int
	CupoRestantePresencial  int
	TotalRecaudoCOP         float64 `json:"total_recaudo_cop"`
	TotalRecaudoUSD         float64 `json:"total_recaudo_usd"`

	TotalPorModalidad     map[string]int
	TotalPorDiaAsistencia map[string]int
	EstadoPorFormaPago    map[string]map[string]int
	IngresosPorFormaPago  map[string]struct {
		TotalCOP int
		TotalUSD int
	}

	TotalSinIglesia     int
	InscripcionesPorDia []InscripcionesDiaDTO
	RecaudoPorModalidad map[string]map[string]float64 `json:"recaudo_por_modalidad"`
}
