// organizar rutas
package routes

import (
	"bubblegum-api/cmd/server/handlers"
	"bubblegum-api/internal/repositories/dulces"
	getdulcebycode "bubblegum-api/internal/usecase/get_dulce_by_code"

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

	//UseCase
	getdulcebycodeUseCase := getdulcebycode.Implementation{
		DulcesProvider: dulceProvider,
	}

	//Handlers
	getdulcebycodeHandler := handlers.GetDulcebyCode{
		UseCase: getdulcebycodeUseCase,
	}

	//endPoint
	p := r.rg.Group("/dulces")

	p.GET("/:codigo", getdulcebycodeHandler.Handle())

}
