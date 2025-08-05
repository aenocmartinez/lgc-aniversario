package dto

type ParticipanteDTO struct {
	Nombre         string `json:"nombre,omitempty"`
	Documento      string `json:"documento,omitempty"`
	Email          string `json:"email,omitempty"`
	Telefono       string `json:"telefono,omitempty"`
	Modalidad      string `json:"modalidad,omitempty"`
	DiasAsistencia string `json:"dias_asistencia,omitempty"`
	Iglesia        string `json:"iglesia,omitempty"`
	Ciudad         string `json:"ciudad,omitempty"`
	HabeasData     bool   `json:"habeas_data,omitempty"`
}
