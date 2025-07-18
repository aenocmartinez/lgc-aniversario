package di

import (
	"lgc/src/dao"
	"lgc/src/domain"
	"lgc/src/infraestructure/database"
	"sync"

	"gorm.io/gorm"
)

type Container struct {
	db       *gorm.DB
	userRepo domain.UserRepository
	// cursoRepo             *dao.CursoDao
	// periodoRepo           *dao.PeriodoDao
	// alumnoRepo            *dao.AlumnoDao
	// cursoPeriodoRepo      *dao.CursoPeriodoDao
	// matriculaRepo         *dao.MatriculaDao
	// maestroRepo           *dao.MaestroDao
	// celebracionRepo       *dao.CelebracionDao
	// grupoRepo             *dao.GrupoDao
	// contenidoTematicoRepo *dao.ContenidoTematicoDao
	// claseRepo             *dao.ClaseDao
}

var (
	instance *Container
	once     sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		db := database.GetDB()
		instance = &Container{
			db:       db,
			userRepo: dao.NewUserDao(db),
			// cursoRepo:             dao.NewCursoDao(db),
			// periodoRepo:           dao.NewPeriodoDao(db),
			// alumnoRepo:            dao.NewAlumnoDao(db),
			// cursoPeriodoRepo:      dao.NewCursoPeriodoDao(db),
			// matriculaRepo:         dao.NewMatriculaDao(db),
			// maestroRepo:           dao.NewMaestroDao(db),
			// celebracionRepo:       dao.NewCelebracionDao(db),
			// grupoRepo:             dao.NewGrupoDao(db),
			// contenidoTematicoRepo: dao.NewContenidoTematicoDao(db),
			// claseRepo:             dao.NewClaseDao(db),
		}
	})
	return instance
}

func (c *Container) GetUserRepository() domain.UserRepository {
	return c.userRepo
}

// func (c *Container) GetCursoRepository() *dao.CursoDao {
// 	return c.cursoRepo
// }

// func (c *Container) GetPeriodoRepository() *dao.PeriodoDao {
// 	return c.periodoRepo
// }

// func (c *Container) GetAlumnoRepository() *dao.AlumnoDao {
// 	return c.alumnoRepo
// }

// func (c *Container) GetCursoPeriodoRepository() *dao.CursoPeriodoDao {
// 	return c.cursoPeriodoRepo
// }

// func (c *Container) GetMatriculaRepository() *dao.MatriculaDao {
// 	return c.matriculaRepo
// }

// func (c *Container) GetMaestroRepository() *dao.MaestroDao {
// 	return c.maestroRepo
// }

// func (c *Container) GetCelebracionRepository() *dao.CelebracionDao {
// 	return c.celebracionRepo
// }

// func (c *Container) GetGrupoRepository() *dao.GrupoDao {
// 	return c.grupoRepo
// }

// func (c *Container) GetContenidoTematicoRepository() *dao.ContenidoTematicoDao {
// 	return c.contenidoTematicoRepo
// }

// func (c *Container) GetClaseRepository() *dao.ClaseDao {
// 	return c.claseRepo
// }
