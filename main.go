// arranca la aplicacion
package main

import (
	"bubblegum-api/cmd/server/routes"
	"bubblegum-api/internal/app/config/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//r = route
	r := gin.Default()

	client := database.Client{}
	db, err := client.Connect()

	if err != nil {
		panic(err)
	}

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	r.Run(":8080")
}

func handlerPing(c *gin.Context) {
	c.String(200, "pong")
}
