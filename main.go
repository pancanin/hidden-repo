package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	auth "questions/auth"
	dals "questions/data"
	handlers "questions/http-handlers"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Secret parsing
	_, authEnabledFlag := os.LookupEnv("AUTH_ENABLED")
	secret := []byte{}

	if authEnabledFlag {
		secretEnvVar, haveSecretEnvVar := os.LookupEnv("JWT_SECRET")

		if haveSecretEnvVar {
			secret = []byte(secretEnvVar)
		} else {
			fmt.Println("You must specify a JWT_SECRET env. variable")
			return
		}
	}

	db, err := gorm.Open(sqlite.Open("questions.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to DB")
	}

	usersDal := dals.NewUsersDal(db)
	questionsDal := dals.NewQuestionsDal(db)
	questionHandler := handlers.NewQuestionHandler(&questionsDal)

	jwtMiddleware := auth.NewJWTAuthMiddleware(secret, &usersDal)
	r := gin.Default()

	r.Use(jwtMiddleware.JWTTokenAuthMiddleware())

	r.POST("/question", questionHandler.Create)
	r.GET("/questions", questionHandler.GetPaginated)
	r.PUT("/question/:id", questionHandler.Update)
	r.DELETE("/question/:id", questionHandler.Delete)

	port := os.Getenv("API_PORT")

	if len(port) == 0 {
		port = "3000"
	}

	portParam := fmt.Sprintf(":%s", port)

	fmt.Printf("Server listening on port %s\n", portParam)

	http.ListenAndServe(portParam, r)
}
