package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Query("token") == "123" {
			log.Println("Auth Success")
			ctx.Next()
		} else {
			ctx.AbortWithStatus(401)
			log.Println("Auth Fails")
		}
	}
}
