package dto

type ParticipanteDTO struct {
	Nombre         string `json:"nombre"`
	Documento      string `json:"documento"`
	Email          string `json:"email"`
	Telefono       string `json:"telefono"`
	Modalidad      string `json:"modalidad"`
	DiasAsistencia string `json:"dias_asistencia"`
	Iglesia        string `json:"iglesia"`
	Ciudad         string `json:"ciudad"`
	HabeasData     bool   `json:"habeas_data"`
}
