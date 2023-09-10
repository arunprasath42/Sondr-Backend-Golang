package middleware

import (
	"errors"
	"fmt"
	"sondr-backend/utils/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TracingMiddleware - middleware
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")
		if value == "" {
			c.AbortWithStatusJSON(400, response.ErrorMessage(400, errors.New("token Not found")))
			return
		}
		tokendata := strings.Split(value, " ") //Ignoring the first value(Bearer keyword)

		/***Validating the secret key***/
		token, err := jwt.Parse(tokendata[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return []byte(viper.GetString("secret.Key")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(400, response.ErrorMessage(400, errors.New("invalid Token")))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("role", claims["role"])
			c.Set("email", claims["email"])
			c.Set("id", claims["id"])
		}
		c.Next()

	}
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")
		fmt.Println("header", value)
		if value == "" {
			c.AbortWithStatusJSON(400, response.ErrorMessage(400, errors.New("token Not found")))
			return
		}
		tokendata := strings.Split(value, " ") //Ignoring the first value(Bearer keyword)
		fmt.Println("token ", tokendata[1])
		/***Validating the secret key***/
		token, err := jwt.Parse(tokendata[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return []byte(viper.GetString("secret.UserKey")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(400, response.ErrorMessage(400, errors.New("invalid Token")))
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("id", claims["id"])
			c.Set("email", claims["email"])
		}
		c.Next()

	}
}
