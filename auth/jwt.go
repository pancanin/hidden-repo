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
	ErrMessages httperrors.ErrorMessages
	Secret      []byte
	UserDal     *dals.UsersDal
}

func NewJWTAuthMiddleware(secret []byte, userDal *dals.UsersDal) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		ErrMessages: httperrors.ErrorMessages{},
		Secret:      secret,
		UserDal:     userDal,
	}
}

func (m *JWTAuthMiddleware) JWTTokenAuthMiddleware() gin.HandlerFunc {
	_, authEnabledFlag := os.LookupEnv("AUTH_ENABLED")

	if authEnabledFlag {
		fmt.Println("Auth enabled!")
		return m.parseAuthReq
	}

	fmt.Println("Auth disabled!")

	return m.noopAuthReq
}

func (m *JWTAuthMiddleware) parseAuthReq(ctx *gin.Context) {
	authHeader := ctx.Request.Header["Authorization"]

	if len(authHeader) == 0 {
		m.ErrMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}

	authHeaderContents := strings.Split(authHeader[0], " ")

	if len(authHeaderContents) < 2 {
		m.ErrMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}

	tokenString := authHeaderContents[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return m.Secret, nil
	})

	if err != nil {
		m.ErrMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Request.Header.Add(models.USER_HEADER_ID, fmt.Sprintf("%s", claims[models.USER_HEADER_ID]))
		ctx.Next()
	} else {
		m.ErrMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}
}

func (m *JWTAuthMiddleware) noopAuthReq(ctx *gin.Context) {
	user, err := m.UserDal.GetByUsername(models.SUPER_USER_NAME)

	if err != nil {
		m.ErrMessages.NotAuthenticated(ctx)
		ctx.Abort()
		return
	}

	fmt.Println("auth user id " + user.ID.String())

	ctx.Request.Header.Add(models.USER_HEADER_ID, user.ID.String())
	ctx.Next()
}
