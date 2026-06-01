// organizar rutas
package routes

import (
	"bubblegum-api/cmd/server/handlers"
	"bubblegum-api/internal/repositories/categorias"
	"bubblegum-api/internal/repositories/dulces"
	"bubblegum-api/internal/repositories/marcas"
	"bubblegum-api/internal/repositories/presentaciones"
	getdulcebycode "bubblegum-api/internal/usecase/get_dulce_by_code"
	getfiltros "bubblegum-api/internal/usecase/get_filtros"

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

	ping := handlers.Ping{}
	r.rg.GET("/ping", ping.Handle())
	//providers
	dulceProvider := dulces.Repository{
		DB: r.db,
	}

	marcaProvider := marcas.Repository{
		DB: r.db,
	}

	categoriaProvider := categorias.Repository{
		DB: r.db,
	}

	presentacionProvider := presentaciones.Repository{
		DB: r.db,
	}

	//UseCase
	getdulcebycodeUseCase := getdulcebycode.Implementation{
		DulcesProvider: dulceProvider,
	}

	getfilros := getfiltros.Implementation{
		MarcasProvider:         marcaProvider,
		CategoriasProvider:     categoriaProvider,
		PresentacionesProvider: presentacionProvider,
	}

	//Handlers
	getdulcebycodeHandler := handlers.GetDulcebyCode{
		UseCase: getdulcebycodeUseCase,
	}

	getfiltrosHandler := handlers.GetFiltros{
		UseCase: getfilros,
	}

	//endPoint
	dulcesGrupo := r.rg.Group("/dulces")
	filtrosGrupo := r.rg.Group("/filtros")

	dulcesGrupo.GET("/:codigo", getdulcebycodeHandler.Handle())
	filtrosGrupo.GET("/", getfiltrosHandler.Handle())
}
