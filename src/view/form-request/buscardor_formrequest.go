package formrequest

type BuscadorInscripcionFormRequest struct {
	Documento string `json:"documento"`
	Nombre    string `json:"nombre"`
	Email     string `json:"email"`
}
