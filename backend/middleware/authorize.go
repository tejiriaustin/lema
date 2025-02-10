package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	constants "github.com/tejiriaustin/lema/constants"
	"github.com/tejiriaustin/lema/env"
	"github.com/tejiriaustin/lema/models"
)

func Authorize(config *env.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Exempt "/auth" routes from authentication
		if strings.HasPrefix(c.Request.URL.Path, "/auth") {
			c.Next()
			return
		}

		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := bearerToken[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return config.GetAsBytes(constants.JwtSecret), nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user := models.AccountInfo{
				Id:       claims["id"].(string),
				FullName: claims["full_name"].(string),
				Email:    claims["email"].(string),
			}

			c.Set("x-user-info", user)
			c.Next()
		} else {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
	}
}
