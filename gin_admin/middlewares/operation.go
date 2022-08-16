package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("日志记录：")
		c.Next()

	}
}
