package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func IsAuth() gin.HandlerFunc {
	return CheckJWT()
}

func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, "")

		if len(bearerToken) == 2 {
			token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte(os.Getenv(JWT_SECRET)), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println(claims["user_id"], claims["user_role"])
			} else {
				c.JSON(422, gin.H{
					"massage": "invalid token"})
				c.Abort()
				return
			}
		} else {

			c.JSON(422, gin.H{
				"massage": "Authorization Token not Provided"})
			c.Abort()
			return
		}
	}

}
