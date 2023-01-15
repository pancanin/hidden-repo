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
	/* DB connection */
	db, err := gorm.Open(sqlite.Open("questions.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to DB")
	}

	/* Dependency wiring */
	usersDal := dals.NewUsersDal(db)
	questionsDal := dals.NewQuestionsDal(db)
	questionHandler := handlers.NewQuestionHandler(&questionsDal)
	jwtMiddleware := auth.NewJWTAuthMiddleware(&usersDal)

	/* Setting up API routes */
	r := gin.Default()

	middlewareFunc, err := jwtMiddleware.JWTTokenAuthMiddleware()

	if err != nil {
		log.Fatalf("There were problems while setting up auth middleware: %s", err)
		return
	}

	r.Use(middlewareFunc)

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

	/* "Fire it up, man!" */
	http.ListenAndServe(portParam, r)
}
