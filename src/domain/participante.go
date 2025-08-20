package domain

import (
	"lgc/src/view/dto"
	"time"
)

type Participante struct {
	id                   int64
	nombre               string
	documento            string
	email                string
	telefono             string
	modalidad            string
	diasAsistencia       string
	iglesia              string
	ciudad               string
	habeasData           bool
	asistenciaConfirmada bool
	qrEnviado            bool
	qrEnviadoAt          *time.Time
	inscripcion          *Inscripcion
	participanteRepo     ParticipanteRepository
}

func NewParticipante(participanteRepo ParticipanteRepository) *Participante {
	return &Participante{
		participanteRepo:     participanteRepo,
		asistenciaConfirmada: false,
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

func (p *Participante) SetAsistenciaConfirmada(asistenciaConfirmada bool) {
	p.asistenciaConfirmada = asistenciaConfirmada
}

func (p *Participante) AsistenciaConfirmada() bool {
	return p.asistenciaConfirmada
}

func (p *Participante) Existe() bool {
	return p.id > 0
}

func (p *Participante) RegistrarAsistencia() bool {
	return p.participanteRepo.RegistrarAsistencia(p.id)
}

func (p *Participante) SetQREnviado(v bool) {
	p.qrEnviado = v
}

func (p *Participante) SetQREnviadoAt(t *time.Time) {
	p.qrEnviadoAt = t
}

func (p *Participante) GetQREnviado() bool {
	return p.qrEnviado
}

func (p *Participante) GetQREnviadoAt() *time.Time {
	return p.qrEnviadoAt
}

func (p *Participante) YaFueEnviadoQR() bool {
	return p.qrEnviado
}

func (p *Participante) ToDTO() dto.ParticipanteDTO {
	return dto.ParticipanteDTO{
		Nombre:               p.GetNombre(),
		Documento:            p.GetDocumento(),
		Email:                p.GetEmail(),
		Telefono:             p.GetTelefono(),
		Modalidad:            p.GetModalidad(),
		DiasAsistencia:       p.GetDiasAsistencia(),
		Iglesia:              p.GetIglesia(),
		Ciudad:               p.GetCiudad(),
		HabeasData:           p.GetHabeasData(),
		AsistenciaConfirmada: p.AsistenciaConfirmada(),
	}
}
