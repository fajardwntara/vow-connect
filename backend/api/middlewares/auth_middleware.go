package middlewares

import (
	"net/http"

	"github.com/fajardwntara/vow-connect/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(helpers.GetEnv("JWT_SECRET", "SECRET_KEY"))

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenStr := ctx.GetHeader("Authorization")

		// If token is none, send the 401 No Authorization response
		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token is required",
			})
			ctx.Abort()
			return
		}

		claims := &jwt.RegisteredClaims{}

		// Parse token and verify the signature using jwtKey
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			ctx.Abort()
			return
		}

		// Save claim "sub" {username} to the context
		ctx.Set("username", claims.Subject)

		// Next to the handler
		ctx.Handler()
	}
}
