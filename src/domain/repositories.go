package domain

import "lgc/src/view/dto"

type UserRepository interface {
	FindByID(id int64) (*User, error)
	FindByEmail(email string) (*User, error)
	Save(user *User) error
	Update(user *User) error
	Delete(id int64) error
}

type InscripcionRepository interface {
	Crear(inscripcion *Inscripcion) bool
	BuscarPorID(inscripcionID int64) *Inscripcion
	BuscarPorDocumento(documento string) *Inscripcion
	ListarInscripcionesPorEstado(estado string) []Inscripcion
	TotalInscripcionesPresenciales() int
	Listar() []Inscripcion
	InscripcionAprobada(inscripcionID int64) bool
	Aprobar(inscripcionID int64) bool
	Anular(inscripcionID int64) bool
	CrearConValidacionDeCupo(inscripcion *Inscripcion, cupoMax int) (bool, error)
}

type EstadisticasRepository interface {
	ObtenerResumenEstadisticas(cupoMax int) dto.EstadisticaResumenDTO
}
