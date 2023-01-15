package auth

import (
	"fmt"
	"os"
	"strings"

	dals "questions/data"
	models "questions/data/models"
	httperrors "questions/http-handlers/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTAuthMiddleware struct {
	errMessages      httperrors.ErrorMessages
	hmacSampleSecret []byte
	userDal          *dals.UsersDal
}

func (m *JWTAuthMiddleware) JWTTokenAuthMiddleware() gin.HandlerFunc {
	m.hmacSampleSecret = []byte("qwertyuiopasdfghjklzxcvbnm123456")
	_, authEnabledFlag := os.LookupEnv("AUTH_ENABLED")

	if authEnabledFlag {
		return m.parseAuthReq
	}

	return m.noopAuthReq
}

func (m *JWTAuthMiddleware) parseAuthReq(ctx *gin.Context) {
	authHeader := ctx.Request.Header["Authorization"]

	if len(authHeader) == 0 {
		m.errMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}

	authHeaderContents := strings.Split(authHeader[0], " ")

	if len(authHeaderContents) < 2 {
		m.errMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}

	tokenString := authHeaderContents[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return m.hmacSampleSecret, nil
	})

	if err != nil {
		m.errMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Request.Header.Add("username", fmt.Sprintf("%s", claims["name"]))
		ctx.Next()
	} else {
		m.errMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}
}

func (m *JWTAuthMiddleware) noopAuthReq(ctx *gin.Context) {
	user, err := m.userDal.GetByUsername(models.SUPER_USER_NAME)

	if err != nil {
		m.errMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}

	ctx.Request.Header.Add(models.USER_HEADER_ID, user.ID.String())
	ctx.Next()
}
