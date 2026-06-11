// organizar rutas
package routes

import (
	getcarritosbyidhandler "bubblegum-api/cmd/server/handlers/get_carrito_by_id"
	getdulcebycodehandler "bubblegum-api/cmd/server/handlers/get_dulce_by_code"
	getfiltroshandler "bubblegum-api/cmd/server/handlers/get_filtros"
	"bubblegum-api/cmd/server/handlers/ping"
	purchasecarritoHandler "bubblegum-api/cmd/server/handlers/purchasecarrito"
	updatecarritohandler "bubblegum-api/cmd/server/handlers/update_carrito"
	ventas "bubblegum-api/internal/repositories/Ventas"
	"bubblegum-api/internal/repositories/carritos"
	"bubblegum-api/internal/repositories/categorias"
	"bubblegum-api/internal/repositories/dulces"
	"bubblegum-api/internal/repositories/marcas"
	"bubblegum-api/internal/repositories/presentaciones"
	"bubblegum-api/internal/repositories/usuarios"
	getcarritobyidusecase "bubblegum-api/internal/usecase/get_carrito_by_id"
	getdulcebycodeusecase "bubblegum-api/internal/usecase/get_dulce_by_code"
	getfiltrosusecase "bubblegum-api/internal/usecase/get_filtros"
	purchasecarritousecase "bubblegum-api/internal/usecase/purchasecarrito"
	updatecarritousecase "bubblegum-api/internal/usecase/update_carrito"

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
	dulcesProvider := dulces.Repository{
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

	carritosProvider := carritos.Repository{
		DB: r.db,
	}

	ventasProvider := ventas.Repository{
		DB: r.db,
	}

	usuariosProvider := usuarios.Repository{
		DB: r.db,
	}
	//UseCase
	getDulceByCodeUseCase := getdulcebycodeusecase.Implementation{
		DulcesProvider:     dulcesProvider,
		CategoriasProvider: &categoriasProvider,
	}

	getfilros := getfiltrosusecase.Implementation{
		MarcasProvider:         marcaProvider,
		CategoriasProvider:     categoriasProvider,
		PresentacionesProvider: presentacionProvider,
	}

	getCarritosByID := getcarritobyidusecase.Implementation{
		CarritoProvider:    carritosProvider,
		DulcesProvider:     dulcesProvider,
		CategoriasProvider: categoriasProvider,
	}

	updateCarrito := updatecarritousecase.Implementation{
		CarritoProvider: carritosProvider,
		DulcesProvider:  dulcesProvider,
	}

	puchaseCarritoUseCase := purchasecarritousecase.Implementation{
		CarritosProvider: carritosProvider,
		UsuariosProvider: usuariosProvider,
		VentasProvider:   ventasProvider,
	}

	//Handlers
	getDulceByCodeHandler := getdulcebycodehandler.GetDulcebyCode{
		UseCase: getDulceByCodeUseCase,
	}

	getFiltrosHandler := getfiltroshandler.GetFiltros{
		UseCase: getfilros,
	}

	getCarritosByIDHandler := getcarritosbyidhandler.GetCarritoByID{
		UseCase: getCarritosByID,
	}

	updateCarritoHandler := updatecarritohandler.UpdateCarrito{
		UseCase: updateCarrito,
	}

	purchaseCarritoH := purchasecarritoHandler.PurchaseCarrito{
		UseCase: puchaseCarritoUseCase,
	}

	// endPoint
	dulcesGrupo := r.rg.Group("/dulces")
	dulcesGrupo.GET("/:codigo", getDulceByCodeHandler.Handle())

	filtrosGrupo := r.rg.Group("/filtros")
	filtrosGrupo.GET("/", getFiltrosHandler.Handle())

	carritosGroup := r.rg.Group("/carritos")
	carritosGroup.GET("/:id", getCarritosByIDHandler.Handle())
	carritosGroup.PUT("/:id/comprar", purchaseCarritoH.Handle())
	carritosGroup.PUT("/:id", updateCarritoHandler.Handle())
}
