package routers

import (
	"net/http"
	. "github.com/soxueren/greenplum-operator/pkg/routers/api"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})


	//actuator
	actuator := router.Group("/actuator")
	actuator.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
		return
	})	

	//swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/v2/api-docs", func(context *gin.Context) {
		context.Redirect(301, "/swagger/doc.json")
	})

	//websocket
	websocket := router.Group("/message")
	websocket.GET("/ws", GetMessage)

	return router
}
