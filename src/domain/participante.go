package domain

import "lgc/src/view/dto"

type Participante struct {
	id               int64
	nombre           string
	documento        string
	email            string
	telefono         string
	modalidad        string
	diasAsistencia   string
	iglesia          string
	ciudad           string
	habeasData       bool
	inscripcion      *Inscripcion
	participanteRepo ParticipanteRepository
}

func NewParticipante(participanteRepo ParticipanteRepository) *Participante {
	return &Participante{
		participanteRepo: participanteRepo,
	}
}

func (p *Participante) SetInscripcion(inscripcion *Inscripcion) {
	p.inscripcion = inscripcion
}

func (p *Participante) GetInscripcion() *Inscripcion {
	return p.inscripcion
}

func (p *Participante) GetID() int64 {
	return p.id
}

func (p *Participante) SetID(id int64) {
	p.id = id
}

func (p *Participante) GetNombre() string {
	return p.nombre
}

func (p *Participante) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Participante) GetDocumento() string {
	return p.documento
}

func (p *Participante) SetDocumento(documento string) {
	p.documento = documento
}

func (p *Participante) GetEmail() string {
	return p.email
}

func (p *Participante) SetEmail(email string) {
	p.email = email
}

func (p *Participante) GetTelefono() string {
	return p.telefono
}

func (p *Participante) SetTelefono(telefono string) {
	p.telefono = telefono
}

func (p *Participante) GetModalidad() string {
	return p.modalidad
}

func (p *Participante) SetModalidad(modalidad string) {
	p.modalidad = modalidad
}

func (p *Participante) GetDiasAsistencia() string {
	return p.diasAsistencia
}

func (p *Participante) SetDiasAsistencia(dias string) {
	p.diasAsistencia = dias
}

func (p *Participante) GetIglesia() string {
	return p.iglesia
}

func (p *Participante) SetIglesia(iglesia string) {
	p.iglesia = iglesia
}

func (p *Participante) GetCiudad() string {
	return p.ciudad
}

func (p *Participante) SetCiudad(ciudad string) {
	p.ciudad = ciudad
}

func (p *Participante) GetHabeasData() bool {
	return p.habeasData
}

func (p *Participante) SetHabeasData(autorizado bool) {
	p.habeasData = autorizado
}

func (p *Participante) Existe() bool {
	return p.id > 0
}

func (p *Participante) ToDTO() dto.ParticipanteDTO {
	return dto.ParticipanteDTO{
		Nombre:         p.GetNombre(),
		Documento:      p.GetDocumento(),
		Email:          p.GetEmail(),
		Telefono:       p.GetTelefono(),
		Modalidad:      p.GetModalidad(),
		DiasAsistencia: p.GetDiasAsistencia(),
		Iglesia:        p.GetIglesia(),
		Ciudad:         p.GetCiudad(),
		HabeasData:     p.GetHabeasData(),
	}
}
