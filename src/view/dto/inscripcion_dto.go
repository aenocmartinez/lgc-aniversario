package dto

type InscripcionDTO struct {
	ID             int64  `json:"id"`
	Nombre         string `json:"nombre"`
	Documento      string `json:"documento"`
	Email          string `json:"email"`
	Telefono       string `json:"telefono"`
	Ciudad         string `json:"ciudad,omitempty"`
	Iglesia        string `json:"iglesia,omitempty"`
	HabeasData     bool   `json:"habeas_data"`
	Estado         string `json:"estado"`
	Revisada       bool   `json:"revisada"`
	Asistencia     string `json:"asistencia"`
	ComprobatePago string `json:"comprobante_pago"`
	FechaRegistro  string `json:"fecha_inscripcion"`
}
