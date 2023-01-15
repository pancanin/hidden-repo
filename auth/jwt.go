package auth

import (
	"fmt"
	"strings"

	httperrors "questions/http-handlers/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTTokenAuthMiddleware() gin.HandlerFunc {
	hmacSampleSecret := []byte("qwertyuiopasdfghjklzxcvbnm123456")
	errMessages := httperrors.ErrorMessages{}

	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header["Authorization"]

		if len(authHeader) == 0 {
			errMessages.NotAuthenticated(ctx)
			ctx.Abort()
			return
		}

		authHeaderContents := strings.Split(authHeader[0], " ")

		if len(authHeaderContents) < 2 {
			errMessages.NotAuthenticated(ctx)
			ctx.Abort()
			return
		}

		tokenString := authHeaderContents[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return hmacSampleSecret, nil
		})

		if err != nil {
			errMessages.NotAuthenticated(ctx)
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Request.Header.Add("username", fmt.Sprintf("%s", claims["name"]))
			ctx.Next()
		} else {
			errMessages.NotAuthenticated(ctx)
			ctx.Abort()
			return
		}
	}

}
