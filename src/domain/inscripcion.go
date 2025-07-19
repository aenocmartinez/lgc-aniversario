package domain

import "lgc/src/view/dto"

type Inscripcion struct {
	id              int64
	nombre          string
	documento       string
	email           string
	telefono        string
	ciudad          string
	iglesia         string
	habeasData      bool
	estado          string
	asistencia      string
	comprobatePago  string
	fechaRegistro   string
	inscripcionRepo InscripcionRepository
}

func NewInscripcion(InscripcionRepo InscripcionRepository) *Inscripcion {
	return &Inscripcion{
		inscripcionRepo: InscripcionRepo,
	}
}

func (f *Inscripcion) SetID(id int64) {
	f.id = id
}

func (f *Inscripcion) GetID() int64 {
	return f.id
}

func (f *Inscripcion) SetNombre(nombre string) {
	f.nombre = nombre
}

func (f *Inscripcion) GetNombre() string {
	return f.nombre
}

func (f *Inscripcion) SetDocumento(documento string) {
	f.documento = documento
}

func (f *Inscripcion) GetDocumento() string {
	return f.documento
}

func (f *Inscripcion) SetEmail(email string) {
	f.email = email
}

func (f *Inscripcion) GetEmail() string {
	return f.email
}

func (f *Inscripcion) SetTelefono(telefono string) {
	f.telefono = telefono
}

func (f *Inscripcion) GetTelefono() string {
	return f.telefono
}

func (f *Inscripcion) SetCiudad(ciudad string) {
	f.ciudad = ciudad
}

func (f *Inscripcion) GetCiudad() string {
	return f.ciudad
}

func (f *Inscripcion) SetIglesia(iglesia string) {
	f.iglesia = iglesia
}

func (f *Inscripcion) GetIglesia() string {
	return f.iglesia
}

func (f *Inscripcion) SetHabeasData(habeasData bool) {
	f.habeasData = habeasData
}

func (f *Inscripcion) GetHabeasData() bool {
	return f.habeasData
}

func (f *Inscripcion) SetEstado(estado string) {
	f.estado = estado
}

func (f *Inscripcion) GetEstado() string {
	return f.estado
}

func (f *Inscripcion) SetAsistencia(asistencia string) {
	f.asistencia = asistencia
}

func (f *Inscripcion) GetAsistencia() string {
	return f.asistencia
}

func (f *Inscripcion) SetComprobantePago(comprobante string) {
	f.comprobatePago = comprobante
}

func (f *Inscripcion) GetComprobantePago() string {
	return f.comprobatePago
}

func (f *Inscripcion) SetFechaRegistro(fecha string) {
	f.fechaRegistro = fecha
}

func (f *Inscripcion) GetFechaRegistro() string {
	return f.fechaRegistro
}

func (f *Inscripcion) Crear() bool {
	return f.inscripcionRepo.Crear(f)
}

func (f *Inscripcion) Existe() bool {
	return f.id > 0
}

func (f *Inscripcion) EstaAprobada() bool {
	return f.estado == "Aprobada"
}

func (f *Inscripcion) EstaPreAprobada() bool {
	return f.estado == "PreAprobada"
}

func (f *Inscripcion) EstaAnulada() bool {
	return f.estado == "Anulada"
}

func (f *Inscripcion) Aprobar() bool {
	return f.inscripcionRepo.Aprobar(f.id)
}

func (f *Inscripcion) Anular() bool {
	return f.inscripcionRepo.Anular(f.id)
}

func (f *Inscripcion) ToDTO() dto.InscripcionDTO {
	return dto.InscripcionDTO{
		ID:             f.id,
		Nombre:         f.nombre,
		Documento:      f.documento,
		Email:          f.email,
		Ciudad:         f.ciudad,
		Iglesia:        f.iglesia,
		HabeasData:     f.habeasData,
		Estado:         f.estado,
		ComprobatePago: f.comprobatePago,
		Telefono:       f.telefono,
		Asistencia:     f.asistencia,
		FechaRegistro:  f.fechaRegistro,
	}
}
