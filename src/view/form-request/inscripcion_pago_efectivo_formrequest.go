package formrequest

type InscripcionPagoEfectivoFormRequest struct {
	Nombre     string `json:"nombre" binding:"required"`
	Documento  string `json:"documento" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Telefono   string `json:"telefono"`
	Iglesia    string `json:"iglesia,omitempty"`
	HabeasData bool   `json:"habeas_data" binding:"required"`
}
