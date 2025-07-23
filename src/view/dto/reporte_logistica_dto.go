package dto

type ReporteLogisticaDTO struct {
	NombreCompleto    string `json:"nombre_completo"`
	NumeroDocumento   string `json:"numero_documento"`
	CorreoElectronico string `json:"correo_electronico"`
	Telefono          string `json:"telefono"`
	DiasAsistencia    string `json:"dias_asistencia"`
}
