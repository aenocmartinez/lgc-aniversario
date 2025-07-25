package main

import (
	"lgc/src/infraestructure/middleware"
	"lgc/src/view/controller"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/assets", "./html/assets")
	r.LoadHTMLGlob("html/*.html")
	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{})
	})

	// Login
	r.POST("/login", controller.Login)

	// Pruebas
	r.GET("/check-db", controller.CheckDBConnection)
	r.GET("/mutant", controller.Mutant)

	// Formulario inscripcion
	r.POST("/realizar-inscripcion", controller.RealizarInscripcion)
	r.POST("/cargar-soporte-pago", controller.CargarArchivoDePago)
	r.GET("/cupos-disponibles", controller.ConsultarCuposDisponibles)
	r.GET("/estadisticas/resumen", controller.ObtenerResumenEstadisticas)
	r.GET("/estadisticas/inscripciones", controller.ListarInscripciones)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/logout", controller.Logout)

		inscripcionGroup := protected.Group("/inscripciones")
		{
			inscripcionGroup.GET("", controller.ListarInscripciones)
			inscripcionGroup.PUT("/rechazar/:id", controller.RechazarInscripcion)
			inscripcionGroup.PUT("/aprobar/:id", controller.AprobarInscripcion)
		}

		reportesGroup := protected.Group("/reportes")
		{
			reportesGroup.GET("/relacion-ingresos", controller.ReporteRelacionDeIngresos)
			reportesGroup.GET("/relacion-ingresos/excel", controller.DescargarRelacionIngresosExcel)
			reportesGroup.GET("/logistica", controller.DescargarReporteLogistica)
		}
	}

	r.Run(":8586")
}
