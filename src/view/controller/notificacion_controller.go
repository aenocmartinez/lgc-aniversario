package controller

import (
	"lgc/src/infraestructure/di"
	"lgc/src/infraestructure/email"
	usecase "lgc/src/usecase/emails"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnviarEmailConQR(c *gin.Context) {

	enviarQR := usecase.NewCorreoEnvioQRUseCase(di.GetContainer().GetParticipanteRepository(), email.NewEmailService(email.GetEmailConfig()))
	err := enviarQR.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Correo enviado"})

}
