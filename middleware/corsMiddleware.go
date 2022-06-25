package middleware

import (
	gin "github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func HandleCrossOriginRequest() gin.HandlerFunc {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://badrri-narayanan.github.io/", "http://localhost:3000"},
	})
	return c
}
