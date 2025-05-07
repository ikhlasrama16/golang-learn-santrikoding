package middlewares

import (
	"example/hello/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			c.Abort()
			return
		}

		// Hapus prefix "Bearer " dari token
		// Header biasanya berbentuk: "Bearer <token>"
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// buat struck untuk menampung klaim token
		claims := &jwt.RegisteredClaims{}

		// parse token dan verifikasi tanda tangan dengan jwtKey
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Kembalikan kunci rahasia untuk memverifikasi token
			return jwtKey, nil
		})

		// Jika token tidak valid atau terjadi error saat parsing
		if err != nil || token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		// Simpan klaim "sub" (username) ke dalam context
		c.Set("username", claims.Subject)

		c.Next()

	}
}
