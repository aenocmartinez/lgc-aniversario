package domain

import "lgc/src/view/dto"

type Inscripcion struct {
	id              int64
	formaPago       string
	montoPagoCOP    int
	montoPagoUSD    int
	urlSoportePago  string
	fechaCreacion   string
	estado          string
	inscripcionRepo InscripcionRepository
}

func NewInscripcion(InscripcionRepo InscripcionRepository) *Inscripcion {
	return &Inscripcion{
		inscripcionRepo: InscripcionRepo,
	}
}

func (i *Inscripcion) SetID(id int64) {
	i.id = id
}

func (i *Inscripcion) GetID() int64 {
	return i.id
}

func (i *Inscripcion) SetFormaPago(formaPago string) {
	i.formaPago = formaPago
}

func (i *Inscripcion) GetFormaPago() string {
	return i.formaPago
}

func (i *Inscripcion) SetMontoPagoCOP(montoPagoCOP int) {
	i.montoPagoCOP = montoPagoCOP
}

func (i *Inscripcion) GetMontoPagoCOP() int {
	return i.montoPagoCOP
}

func (i *Inscripcion) SetMontoPagoUSD(montoPagoUSD int) {
	i.montoPagoUSD = montoPagoUSD
}

func (i *Inscripcion) GetMontoPagoUSD() int {
	return i.montoPagoUSD
}

func (i *Inscripcion) SetUrlSoportePago(urlSoportePago string) {
	i.urlSoportePago = urlSoportePago
}

func (i *Inscripcion) GetUrlSoportePago() string {
	return i.urlSoportePago
}

func (i *Inscripcion) SetEstado(estado string) {
	i.estado = estado
}

func (i *Inscripcion) GetEstado() string {
	return i.estado
}

func (i *Inscripcion) SetFechaCreacion(fechaCreacion string) {
	i.fechaCreacion = fechaCreacion
}

func (i *Inscripcion) GetFechaCreacion() string {
	return i.fechaCreacion
}

func (i *Inscripcion) AgregarParticipante(participante Participante) bool {
	return i.inscripcionRepo.AgregarParticipante(i.id, participante)
}

func (i *Inscripcion) Participantes() []Participante {
	return i.inscripcionRepo.ObtenerParticipantes(i.id)
}

func (i *Inscripcion) Crear() bool {
	return i.inscripcionRepo.Crear(i)
}

func (i *Inscripcion) Existe() bool {
	return i.id > 0
}

func (i *Inscripcion) EstaAprobada() bool {
	return i.estado == "Aprobada"
}

func (i *Inscripcion) EstaPreAprobada() bool {
	return i.estado == "PreAprobada"
}

func (i *Inscripcion) EstaRechazada() bool {
	return i.estado == "Rechazada"
}

func (i *Inscripcion) Aprobar() bool {
	return i.inscripcionRepo.Aprobar(i.id)
}

func (i *Inscripcion) Anular() bool {
	return i.inscripcionRepo.Rechazar(i.id)
}

func (i *Inscripcion) ToDTO() dto.InscripcionDTO {
	return dto.InscripcionDTO{
		ID:             i.id,
		FormaPago:      i.formaPago,
		MontoPagoCOP:   i.montoPagoCOP,
		MontoPagoUSD:   i.montoPagoUSD,
		UrlSoportePago: i.urlSoportePago,
		FechaCreacion:  i.fechaCreacion,
		Estado:         i.estado,
	}
}
