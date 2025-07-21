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
	CrearConValidacionDeCupo(inscripcion *Inscripcion, participantes []Participante, cupoMax int) error
	Listar() []Inscripcion
	BuscarPorID(inscripcionID int64) Inscripcion
	AgregarParticipante(inscripcionID int64, participante Participante) bool
	ObtenerParticipantes(inscripcionID int64) []Participante
	Aprobar(inscripcionID int64) bool
	Rechazar(inscripcionID int64) bool
	CuposDisponibles(cupoMax int) (ocupados int, disponibles int)
}

type ParticipanteRepository interface {
	CrearParticipante(participante Participante) bool
	BuscarPartipante(participanteID int64) Participante
}

type EstadisticasRepository interface {
	ObtenerResumenEstadisticas(cupoMax int) dto.EstadisticaResumenDTO
}
