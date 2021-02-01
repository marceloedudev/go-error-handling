package middleware

import (
	"fmt"
	"go-error-handling/exceptions"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *exceptions.HttpException:
					{
						errors := err.(*exceptions.HttpException)
						errors.Error = http.StatusText(errors.Status)
						errors.Timestamp = time.Now()
						errors.Path = c.Request.URL.Path
						c.JSON(errors.Status, errors)
						return
					}
				default:
					{
						c.JSON(http.StatusInternalServerError, &exceptions.HttpException{
							Message:   fmt.Sprintf("%v", err),
							Status:    http.StatusInternalServerError,
							Error:     http.StatusText(http.StatusInternalServerError),
							Causes:    nil,
							Timestamp: time.Now(),
							Path:      c.Request.URL.Path,
						})
						return
					}
				}
			}
		}()

		c.Next()
	}
}
