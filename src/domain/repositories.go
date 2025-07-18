package domain

type UserRepository interface {
	FindByID(id int64) (*User, error)
	FindByEmail(email string) (*User, error)
	Save(user *User) error
	Update(user *User) error
	Delete(id int64) error
}

type InscripcionRepository interface {
	Listar() []Inscripcion
	ListarInscripcionesPorEstado(estado string) []Inscripcion
	Crear(inscripcion *Inscripcion) bool
	BuscarPorID(inscripcionID int64) *Inscripcion
	BuscarPorDocumento(documento string) *Inscripcion
	InscripcionValidada(inscripcionID int64) bool
	Validar(inscripcionID int64, validado bool) bool
}
