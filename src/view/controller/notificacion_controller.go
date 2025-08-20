package controller

import (
	"lgc/src/infraestructure/di"
	"lgc/src/infraestructure/email"
	usecase "lgc/src/usecase/emails"
	usecaseParticipante "lgc/src/usecase/participantes"
	formrequest "lgc/src/view/form-request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnviarEmailConQR(c *gin.Context) {

	enviarQR := usecase.NewCorreoEnvioQRUseCase(di.GetContainer().GetParticipanteRepository(), email.NewEmailService(email.GetEmailConfig()))
	enviarQR.BatchSize = 50
	enviarQR.Workers = 3
	err := enviarQR.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Correo enviado"})

}

func BuscarParticipante(c *gin.Context) {
	var documento string = c.Query("documento")

	if documento == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'documento' es requerido"})
		return
	}

	buscarParticipante := usecaseParticipante.NewBuscarParticipantePorDocumentoUseCase(di.GetContainer().GetParticipanteRepository())
	response := buscarParticipante.Execute(documento)

	c.JSON(response.StatusCode, response)
}

func ConfirmarAsistencia(c *gin.Context) {
	var req formrequest.ConfirmarAsistenciaFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	confirmarAsistencia := usecaseParticipante.NewConfirmaAsistenciaUseCase(
		di.GetContainer().GetParticipanteRepository(),
	)

	response := confirmarAsistencia.Execute(req.Documento)

	c.JSON(response.StatusCode, response)
}

func VisualizarParticipante(c *gin.Context) {
	c.HTML(http.StatusOK, "visualizar.html", nil)
}
