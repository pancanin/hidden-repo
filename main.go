package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	dals "questions/data"
	handlers "questions/http-handlers"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("questions.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to DB")
	}

	questionsDal := dals.NewQuestionsDal(db)
	questionHandler := handlers.NewQuestionHandler(&questionsDal)

	r := gin.Default()

	tokenString := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NzM3OTEwODcsImV4cCI6MTcwNTMyNzA4NywiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.6CPbWC2cIfxp-RyB7JqzRnxNGOg1ia3qP2pWWyr1RgY"
	hmacSampleSecret := []byte("qwertyuiopasdfghjklzxcvbnm123456")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["GivenName"], claims["Email"])
	} else {
		fmt.Println(err)
	}

	r.POST("/question", questionHandler.Create)
	r.GET("/questions", questionHandler.GetAll)
	r.PUT("/question/:id", questionHandler.Update)
	r.DELETE("/question/:id", questionHandler.Delete)

	port := os.Getenv("API_PORT")

	if len(port) == 0 {
		port = "3000"
	}

	portParam := fmt.Sprintf(":%s", port)

	fmt.Printf("Server listening on port %s", portParam)

	http.ListenAndServe(portParam, r)
}
