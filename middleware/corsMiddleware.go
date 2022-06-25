package middleware

import (
	gin "github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func HandleCrossOriginRequest() gin.HandlerFunc {
	c := cors.Default()
	return c
}
