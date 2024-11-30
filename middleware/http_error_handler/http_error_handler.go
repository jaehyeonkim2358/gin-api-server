package httperrorhandler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				if err, ok := e.Err.(interface{ StackTrace() errors.StackTrace }); ok {
					for _, frame := range err.StackTrace() {
						log.Printf("%+s:%d\n", frame, frame)
					}
				}
    	}

			c.IndentedJSON(c.Writer.Status(), gin.H{"message": c.Errors[0].Error()})
		}
	}
}
