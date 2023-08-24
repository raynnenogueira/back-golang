package main // pacote principal

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// instancia de rotas do GIN
	r := gin.Default()

	// rota para o endpoint /ping
	r.GET("/ping", func(c *gin.Context) { //chamando uma rota GET (buscar)
		c.JSON(200, gin.H {
			"message":"pong",
		})
	}) 

	r.GET("/health", func(c *gin.Context) { //chamando uma rota GET (buscar)
		c.JSON(200, gin.H {
			"message":"aplication is running.",
		})
	}) 

	//porta que o servidor ira executar
	r.Run(":8080")
}