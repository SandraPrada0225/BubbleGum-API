// organizar rutas
package routes

import (
	getdulcebycodehandler "bubblegum-api/cmd/server/handlers/get_dulce_by_code"
	getfiltroshandler "bubblegum-api/cmd/server/handlers/get_filtros"
	"bubblegum-api/cmd/server/handlers/ping"
	"bubblegum-api/internal/repositories/categorias"
	"bubblegum-api/internal/repositories/dulces"
	"bubblegum-api/internal/repositories/marcas"
	"bubblegum-api/internal/repositories/presentaciones"
	getdulcebycodeusecase "bubblegum-api/internal/usecase/get_dulce_by_code"
	getfiltrosusecase "bubblegum-api/internal/usecase/get_filtros"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct{}

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *gorm.DB
}

func NewRouter(eng *gin.Engine, db *gorm.DB) Router {
	return &router{
		eng: eng,
		db:  db,
	}
}

func (r router) MapRoutes() {
	r.rg = r.eng.Group("/api")

	pingHandler := ping.Ping{}
	r.rg.GET("/ping", pingHandler.Handle())
	//providers
	dulceProvider := dulces.Repository{
		DB: r.db,
	}

	marcaProvider := marcas.Repository{
		DB: r.db,
	}

	categoriasProvider := categorias.Repository{
		DB: r.db,
	}

	presentacionProvider := presentaciones.Repository{
		DB: r.db,
	}

	//UseCase
	getDulceByCodeUseCase := getdulcebycodeusecase.Implementation{
		DulcesProvider:     dulceProvider,
		CategoriasProvider: &categoriasProvider,
	}

	getfilros := getfiltrosusecase.Implementation{
		MarcasProvider:         marcaProvider,
		CategoriasProvider:     categoriasProvider,
		PresentacionesProvider: presentacionProvider,
	}

	//Handlers
	getDulceByCodeHandler := getdulcebycodehandler.GetDulcebyCode{
		UseCase: getDulceByCodeUseCase,
	}

	getFiltrosHandler := getfiltroshandler.GetFiltros{
		UseCase: getfilros,
	}

	// endPoint
	dulcesGrupo := r.rg.Group("/dulces")
	dulcesGrupo.GET("/:codigo", getDulceByCodeHandler.Handle())

	filtrosGrupo := r.rg.Group("/filtros")
	filtrosGrupo.GET("/", getFiltrosHandler.Handle())
}
