package routes

import (
	"fmt"
	"go-error-handling/dtos"
	"go-error-handling/exceptions"
	"go-error-handling/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AddRoutes() *gin.Engine {
	app := gin.Default()

	app.Use(middleware.HandleErrors())

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	app.POST("/status", func(c *gin.Context) {
		var formData dtos.StatusJSON

		if err := c.ShouldBindJSON(&formData); err != nil {
			panic(exceptions.NewInternalServerError("There was a problem with validation", err))
		}

		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	app.POST("/panic1", func(c *gin.Context) {
		panic("test")
	})

	app.POST("/panic2", func(c *gin.Context) {
		_, err := os.Create("/tmp/file")
		if err != nil {
			panic(err)
		}
	})

	app.NoRoute(func(c *gin.Context) {
		panic(exceptions.NewNotFoundError(fmt.Sprintf("Route '%s' was not found", c.Request.URL.Path)))
	})

	return app
}
