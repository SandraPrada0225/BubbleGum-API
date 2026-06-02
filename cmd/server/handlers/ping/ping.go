// maneja peticiones HTTP
package ping

import "github.com/gin-gonic/gin"

type Ping struct {
}

// el metodo pertenece al tipo ping
func (handler Ping) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "Pong")
	}
}
