package di

import (
	"lgc/src/dao"
	"lgc/src/domain"
	"lgc/src/infraestructure/database"
	"sync"

	"gorm.io/gorm"
)

type Container struct {
	db               *gorm.DB
	userRepo         domain.UserRepository
	inscripcionRepo  domain.InscripcionRepository
	estadisticasRepo domain.EstadisticasRepository
}

var (
	instance *Container
	once     sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		db := database.GetDB()
		instance = &Container{
			db:               db,
			userRepo:         dao.NewUserDao(db),
			inscripcionRepo:  dao.NewInscripcionDao(db),
			estadisticasRepo: dao.NewEstadisticasDao(db),
		}
	})
	return instance
}

func (c *Container) GetUserRepository() domain.UserRepository {
	return c.userRepo
}

func (c *Container) GetInscripcionRepository() domain.InscripcionRepository {
	return c.inscripcionRepo
}

func (c *Container) GetEstadisticasRepository() domain.EstadisticasRepository {
	return c.estadisticasRepo
}
